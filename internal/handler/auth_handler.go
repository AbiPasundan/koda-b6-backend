package handler

import (
	"backend/internal/helper"
	"backend/internal/models"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func NewAuthHandler(repo *service.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: repo,
	}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var password models.AuthLogin

	ctx.ShouldBindJSON(&password)
}

func (h *AuthHandler) Register(ctx *gin.Context) {

	var req models.AuthRegister

	err := ctx.ShouldBindJSON(&req)
	helper.BadRequest(ctx, "Invalid request body", nil, err)

	h.AuthService.Register(&models.AuthRegister{})

	ctx.JSON(200, gin.H{"message": "user created"})
}
