package routes

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, h *handler.AuthHandler) {
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
}
