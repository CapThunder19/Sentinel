package repository

import (
	"context"

	"github.com/CapThunder19/Sentinel/backend/auth-service/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	GetByID(id uuid.UUID) (*models.User, error)
}

type PostgreSQLUserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *PostgreSQLUserRepository {
	return &PostgreSQLUserRepository{
		db: db,
	}
}

func (r *PostgreSQLUserRepository) Create(user *models.User) error {
	query := `
	INSERT INTO users (
		username,
		email,
		password_hash,
		role,
		is_verified
	)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, created_at, updated_at;
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.Role,
		user.IsVerified,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *PostgreSQLUserRepository) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}

	query := `
	SELECT
		id,
		username,
		email,
		password_hash,
		role,
		is_verified,
		created_at,
		updated_at
	FROM users
	WHERE email = $1
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostgreSQLUserRepository) GetByID(id string) (*models.User, error) {

	user := &models.User{}

	query := `
	SELECT
		id,
		username,
		email,
		password_hash,
		role,
		is_verified,
		created_at,
		updated_at
	FROM users
	WHERE id = $1
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
