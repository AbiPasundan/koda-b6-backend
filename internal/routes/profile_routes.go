package routes

import (
	"backend/internal/handler"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(r *gin.Engine, h *handler.ProfileHandler) {
	profile := r.Group("/profile")
	profile.Use(middleware.JWTMiddleware(), middleware.RoleMiddleware("admin", "user"))
	{
		profile.GET("", h.GetMyProfile)
		profile.POST("/update-profile", h.UpdateProfile)
	}
}
