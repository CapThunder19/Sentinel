package main

import (

    "github.com/gin-gonic/gin"
     
	"github.com/CapThunder19/Sentinel/backend/auth-service/internal/routes"
    "github.com/CapThunder19/Sentinel/backend/shared/config"
    "github.com/CapThunder19/Sentinel/backend/shared/logger"
)

func main() {
	cfg := config.Load()
	logger.Init()
	logger.Info("Starting Auth Service...")

	router := gin.Default()

	routes.RegisterRoutes(router)

	logger.Info("Auth Service listening on port " + cfg.Port)

	if err := router.Run(":" + cfg.Port); err != nil {
		logger.Error(err.Error())
	}

	

}