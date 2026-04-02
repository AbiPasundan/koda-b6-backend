package routes

import (
	"backend/internal/handler"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(r *gin.Engine, h *handler.ProfileHandler) {
	admin := r.Group("/admin")
	admin.Use(middleware.JWTMiddleware(), middleware.RoleMiddleware("admin", "user"))
	admin.GET("/profile", h.GetMyProfile)
}
