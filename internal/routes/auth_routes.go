package routes

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, h *handler.AuthHandler) {
	r.POST("/auth/register", h.Register)
	r.POST("/auth/login", h.Login)
	r.POST("/forgot-password", h.ResetPassword)
	r.POST("/request-forgot-password", h.RequestForgotPassword)
}
