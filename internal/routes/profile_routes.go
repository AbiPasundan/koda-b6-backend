package routes

import (
	"backend/internal/handler"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(r *gin.Engine, h *handler.ProfileHandler) {
	r.Use(middleware.JWTMiddleware(), middleware.RoleMiddleware("admin", "user"))
	{
		r.GET("/profile", h.GetMyProfile)
		r.PATCH("/update-profile", h.UpdateProfile)
	}
}
