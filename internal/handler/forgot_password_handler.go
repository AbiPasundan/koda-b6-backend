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
	_, err := f.ForgotPasswordService.RequestForgotPassword("email")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "invalid requset " + err.Error(),
			Results: nil,
		})
	}

}
