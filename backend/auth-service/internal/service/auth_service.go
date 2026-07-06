package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/CapThunder19/Sentinel/backend/auth-service/internal/models"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
}

type AuthService struct {
	repo UserRepository
}

func NewAuthService(repo UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
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
