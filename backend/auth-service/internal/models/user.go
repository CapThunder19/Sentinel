package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID
	Username     string
	Email        string
	PasswordHash string
	Role         string
	IsVerified   bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
