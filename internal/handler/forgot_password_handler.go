package handler

import (
	"backend/internal/models"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ForgotPasswordHandler struct {
	ForgotPasswordService *service.ForgotPasswordService
}

func NewForgotPasswordHandler(service *service.ForgotPasswordService) *ForgotPasswordHandler {
	return &ForgotPasswordHandler{
		ForgotPasswordService: service,
	}
}

func (f *ForgotPasswordHandler) RequestForgotPassword(ctx *gin.Context) {
	var email models.JustEmail

	if err := ctx.ShouldBindJSON(&email); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid Email request body: " + err.Error(),
			Results: nil,
		})
		return
	}

	code, err := f.ForgotPasswordService.RequestForgotPassword(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to request forgot password: " + err.Error(),
			Results: code,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "OTP Generated",
		Results: code,
	})
}

func (f *ForgotPasswordHandler) ResetPassword(ctx *gin.Context) {
	var req models.ResetPasswordInput

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid request body: " + err.Error(),
			Results: nil,
		})
		return
	}

	if err := f.ForgotPasswordService.ResetPassword(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to reset password: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Password has been reset successfully",
	})
}
