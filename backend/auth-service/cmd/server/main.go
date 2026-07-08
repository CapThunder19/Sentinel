package main

import (
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	"github.com/CapThunder19/Sentinel/backend/auth-service/internal/handler"
	"github.com/CapThunder19/Sentinel/backend/auth-service/internal/repository"
	"github.com/CapThunder19/Sentinel/backend/auth-service/internal/routes"
	"github.com/CapThunder19/Sentinel/backend/auth-service/internal/service"
	"github.com/CapThunder19/Sentinel/backend/shared/config"
	"github.com/CapThunder19/Sentinel/backend/shared/database"
	"github.com/CapThunder19/Sentinel/backend/shared/logger"
)

func main() {
	_ = godotenv.Load(".env")

	cfg := config.Load()
	cfg.Validate()
	logger.Init()
	logger.Info("Starting Auth Service...")

	pool, err := database.Connect(cfg)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer pool.Close()

	router := gin.Default()

	repo := repository.NewUserRepository(pool)

	authService := service.NewAuthService(
		repo,
		cfg.JWTSecret,
	)

	authHandler := handler.NewAuthHandler(authService)

	routes.RegisterRoutes(
		router,
		authHandler,
		cfg.JWTSecret,
	)

	logger.Info("Auth Service listening on port " + cfg.Port)

	if err := router.Run(":" + cfg.Port); err != nil {
		logger.Error(err.Error())
	}

}
