package routes

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, h *handler.ForgotPasswordHandler) {
	r.GET("/auth/forgotpassword", h.RequestForgotPassword)
}
