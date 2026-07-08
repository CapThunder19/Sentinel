package handler

import (
	"net/http"

	"github.com/CapThunder19/Sentinel/backend/auth-service/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {

	var req service.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.authService.Register(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *AuthHandler) Login(c *gin.Context) {

	var req service.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.authService.Login(req)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *AuthHandler) Me(c *gin.Context) {

	userID := c.GetString("userID")
	email := c.GetString("email")

	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"email":   email,
	})
}
