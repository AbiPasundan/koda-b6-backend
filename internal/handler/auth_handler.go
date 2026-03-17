package handler

import (
	"backend/internal/helper"
	"backend/internal/models"
	"backend/internal/service"
	"fmt"
	"net/http"

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
	var user models.AuthLogin

	err := ctx.ShouldBindJSON(&user)
	if helper.BadRequest(ctx, "Invalid request body", nil, err) {
		return
	}

	users, err := h.AuthService.FindEmail(user.Email)
	if helper.CustomeError(ctx, http.StatusUnauthorized, "Unauthorized", users, err) {
		return
	}

	fmt.Println(users)

	helper.ResponseOk(ctx, "Success Login", users)
}

func (h *AuthHandler) Register(ctx *gin.Context) {

	var req models.AuthRegister

	err := ctx.ShouldBindJSON(&req)
	if helper.BadRequest(ctx, "Invalid request body", nil, err) {
		return
	}

	h.AuthService.Register(&models.AuthRegister{})

	helper.ResponseOk(ctx, "Success Create User", nil)
}
