package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/CapThunder19/Sentinel/backend/auth-service/internal/handler"
	"github.com/CapThunder19/Sentinel/backend/shared/auth"
)

func RegisterRoutes(
	router *gin.Engine,
	authHandler *handler.AuthHandler,
	jwtSecret string,
) {
	router.GET("/", handler.Health)

	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	protected := router.Group("/")
	protected.Use(auth.JWTMiddleware(jwtSecret))

	protected.GET("/me", authHandler.Me)
}
