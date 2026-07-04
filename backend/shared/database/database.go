package database

import (
	"context"

	"github.com/CapThunder19/Sentinel/backend/shared/config"
	"github.com/CapThunder19/Sentinel/backend/shared/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(cfg *config.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)

	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, err
	}

	if err != nil {
		pool.Close()
		return nil, err
	}

	logger.Info("Connected to PostgreSQL")

	return pool, nil
}
