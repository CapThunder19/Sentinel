package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/CapThunder19/Sentinel/backend/auth-service/internal/models"
	"github.com/CapThunder19/Sentinel/backend/shared/auth"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	GetByID(id string) (*models.User, error)
}

type AuthService struct {
	repo      UserRepository
	jwtSecret string
}

func NewAuthService(
	repo UserRepository,
	jwtSecret string,
) *AuthService {

	return &AuthService{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) Register(req RegisterRequest) (*RegisterResponse, error) {

	existingUser, err := s.repo.GetByEmail(req.Email)

	if err == nil && existingUser != nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Role:         "USER",
		IsVerified:   false,
	}

	err = s.repo.Create(user)

	if err != nil {
		return nil, err
	}

	return &RegisterResponse{
		ID:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (s *AuthService) Login(req LoginRequest) (*LoginResponse, error) {

	user, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)

	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := auth.GenerateToken(
		user.ID.String(),
		user.Email,
		s.jwtSecret,
	)

	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
	}, nil
}

func (s *AuthService) GetProfile(userID string) (*models.User, error) {
	return s.repo.GetByID(userID)
}
