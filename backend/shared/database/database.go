package database

import (
	"context"
	"time"

	"github.com/CapThunder19/Sentinel/backend/shared/config"
	"github.com/CapThunder19/Sentinel/backend/shared/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(cfg *config.Config) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if poolConfig.ConnConfig.ConnectTimeout == 0 {
		poolConfig.ConnConfig.ConnectTimeout = 5 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)

	if err != nil {
		return nil, err
	}

	pingCtx, pingCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer pingCancel()

	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, err
	}

	logger.Info("Connected to PostgreSQL")

	return pool, nil
}
