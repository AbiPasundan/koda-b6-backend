package handler

import (
	"backend/internal/helper"
	"backend/internal/middleware"
	"backend/internal/models"
	"backend/internal/service"
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
	// var user models.User

	// err := ctx.ShouldBindJSON(&user)
	// if helper.BadRequest(ctx, "Invalid request body", nil, err) {
	// 	return
	// }

	// users, err := h.AuthService.FindEmail(user.Email)
	// if helper.CustomeError(ctx, http.StatusUnauthorized, "Unauthorized", users, err) {
	// 	return
	// }

	// fmt.Println(users)

	// helper.ResponseOk(ctx, "Success Login", users)

	var user models.AuthLogin

	err := ctx.ShouldBindJSON(&user)
	if helper.BadRequest(ctx, "Invalid request body", nil, err) {
		return
	}

	h.AuthService.FindEmail(user.Email)

	token, err := middleware.GenerateToken(user.Id)
	if err != nil {
		helper.CustomeError(ctx, http.StatusInternalServerError, "Failed generate token", nil, err)
		return
	}

	helper.ResponseOk(ctx, "Success Login", token)

}

func (h *AuthHandler) Register(ctx *gin.Context) {

	var req models.AuthRegister

	err := ctx.ShouldBindJSON(&req)
	if helper.BadRequest(ctx, "Invalid request body", nil, err) {
		return
	}

	h.AuthService.Register(&req)

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
