package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/joho/godotenv"

	"github.com/CapThunder19/Sentinel/backend/auth-service/internal/models"
	"github.com/CapThunder19/Sentinel/backend/shared/config"
	"github.com/CapThunder19/Sentinel/backend/shared/database"

	"github.com/CapThunder19/Sentinel/backend/shared/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	logger.Init()

	_ = godotenv.Load("../../../.env.test")

	cfg := config.Load()
	t.Log("DATABASE_URL:", cfg.DatabaseURL)
	pool, err := database.Connect(cfg)
	require.NoError(t, err)

	defer pool.Close()

	repo := NewUserRepository(pool)

	user := &models.User{
		Username:     "anirudh",
		Email:        "ani@example.com",
		PasswordHash: "hashed-password",
		Role:         "USER",
		IsVerified:   false,
	}

	err = repo.Create(user)

	require.NoError(t, err)

	assert.NotEqual(t, user.ID.String(), "")
	assert.False(t, user.CreatedAt.IsZero())
	assert.False(t, user.UpdatedAt.IsZero())

	_, err = pool.Exec(
		context.Background(),
		"DELETE FROM users WHERE email=$1",
		user.Email,
	)

	require.NoError(t, err)
}

func TestGetUserByEmail(t *testing.T) {
	logger.Init()

	err := godotenv.Load("../../../.env.test")
	require.NoError(t, err)

	cfg := config.Load()

	pool, err := database.Connect(cfg)
	require.NoError(t, err)

	defer pool.Close()

	repo := NewUserRepository(pool)

	user := &models.User{
		Username:     "anirudh",
		Email:        fmt.Sprintf("ani_%d@example.com", time.Now().UnixNano()),
		PasswordHash: "hashed-password",
		Role:         "USER",
		IsVerified:   false,
	}

	err = repo.Create(user)
	require.NoError(t, err)

	foundUser, err := repo.GetByEmail(user.Email)

	require.NoError(t, err)
	require.NotNil(t, foundUser)

	assert.Equal(t, user.Username, foundUser.Username)
	assert.Equal(t, user.Email, foundUser.Email)
	assert.Equal(t, user.PasswordHash, foundUser.PasswordHash)
	assert.Equal(t, user.Role, foundUser.Role)
	assert.Equal(t, user.IsVerified, foundUser.IsVerified)

	_, err = pool.Exec(
		context.Background(),
		"DELETE FROM users WHERE email=$1",
		user.Email,
	)

	require.NoError(t, err)
}
