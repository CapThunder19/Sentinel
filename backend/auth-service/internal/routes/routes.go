package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/CapThunder19/Sentinel/backend/auth-service/internal/handler"
)

func RegisterRoutes(
	router *gin.Engine,
	authHandler *handler.AuthHandler,
) {
	router.GET("/", handler.Health)

	router.POST("/register", authHandler.Register)
}
