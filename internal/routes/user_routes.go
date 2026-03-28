package routes

import (
	"backend/internal/handler"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, h *handler.UserHandler) {
	admin := r.Group("/admin")
	admin.Use(middleware.JWTMiddleware(), middleware.RoleMiddleware("admin"))
	admin.GET("/users", h.Home)
	admin.GET("/users/:id", h.GetUserById)
	admin.POST("/users", h.AddUser)
	admin.PATCH("/users/:id", h.UpdateUser)
	admin.DELETE("/users/:id", h.DeleteUser)
}
