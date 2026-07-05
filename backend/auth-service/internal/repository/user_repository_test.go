package repository

import (
	"context"
	"testing"

	"github.com/joho/godotenv"

	"github.com/CapThunder19/Sentinel/backend/auth-service/internal/models"
	"github.com/CapThunder19/Sentinel/backend/shared/config"
	"github.com/CapThunder19/Sentinel/backend/shared/database"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/CapThunder19/Sentinel/backend/shared/logger"

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
