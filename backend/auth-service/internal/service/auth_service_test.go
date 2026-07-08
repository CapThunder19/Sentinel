package service

import (
	"errors"
	"testing"

	"github.com/CapThunder19/Sentinel/backend/auth-service/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type FakeUserRepository struct {
	users map[string]*models.User
}

func NewFakeUserRepository() *FakeUserRepository {
	return &FakeUserRepository{
		users: make(map[string]*models.User),
	}
}

func (r *FakeUserRepository) Create(user *models.User) error {
	r.users[user.Email] = user
	return nil
}

func (r *FakeUserRepository) GetByEmail(email string) (*models.User, error) {
	user, ok := r.users[email]

	if !ok {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func TestRegisterSuccess(t *testing.T) {
	repo := NewFakeUserRepository()

	service := NewAuthService(repo, "test-secret")

	req := RegisterRequest{
		Username: "anirudh",
		Email:    "ani@test.com",
		Password: "password123",
	}

	res, err := service.Register(req)

	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Equal(t, req.Username, res.Username)
	assert.Equal(t, req.Email, res.Email)
}

func TestLoginSuccess(t *testing.T) {

	repo := NewFakeUserRepository()

	service := NewAuthService(
		repo,
		"test-secret",
	)

	registerReq := RegisterRequest{
		Username: "anirudh",
		Email:    "ani@test.com",
		Password: "password123",
	}

	_, err := service.Register(registerReq)
	require.NoError(t, err)

	loginReq := LoginRequest{
		Email:    "ani@test.com",
		Password: "password123",
	}

	res, err := service.Login(loginReq)

	require.NoError(t, err)
	require.NotNil(t, res)

	assert.NotEmpty(t, res.AccessToken)
	assert.Equal(t, "Bearer", res.TokenType)
}
