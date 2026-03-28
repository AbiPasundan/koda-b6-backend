package handler

import (
	"backend/internal/helper"
	"backend/internal/middleware"
	"backend/internal/models"
	"backend/internal/service"
	"net/http"
	"strings"

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
	var req models.AuthLogin

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.BadRequest(ctx, "Invalid request", nil, err)
		return
	}

	user, err := h.AuthService.Login(req.Email, req.Password)
	if err != nil {
		helper.BadRequest(ctx, "Wrong Email or Password", nil, err)
		return
	}

	token, err := middleware.GenerateToken(user.Id, user.Role)
	if err != nil {
		helper.CustomeError(ctx, http.StatusInternalServerError, "Failed generate token", nil, err)
		return
	}

	helper.ResponseOk(ctx, "Success Login", token)
}

func (h *AuthHandler) Register(ctx *gin.Context) {

	var req models.AuthRegister

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.BadRequest(ctx, "Invalid request body", nil, err)
		return
	}

	err := h.AuthService.Register(&req)
	if err != nil {

		// handle duplicate email
		if strings.Contains(err.Error(), "duplicate key") {
			helper.BadRequest(ctx, "Email already exists", nil, err)
			return
		}

		helper.CustomeError(ctx, http.StatusInternalServerError, "Failed register", nil, err)
		return
	}

	helper.ResponseOk(ctx, "Success Create User", nil)
}

func (h *AuthHandler) RequestForgotPassword(ctx *gin.Context) {
	var req models.AuthForgotPassword
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.BadRequest(ctx, "Invalid request body", nil, err)
		return
	}
	err := h.AuthService.ForgotPasswordRequest(&req)
	if err != nil {
		helper.BadRequest(ctx, err.Error(), nil, err)
		return
	}
	helper.ResponseOk(ctx, "OTP sent", nil)
}

func (h *AuthHandler) ResetPassword(ctx *gin.Context) {
	var req models.ResetPasswordRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.BadRequest(ctx, "Invalid request", nil, err)
		return
	}

	err := h.AuthService.ResetPassword(req)
	if err != nil {
		helper.BadRequest(ctx, err.Error(), nil, err)
		return
	}

	helper.ResponseOk(ctx, "Password reset success", nil)
}
