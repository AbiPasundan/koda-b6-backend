package routes

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, h *handler.UserHandler) {
	r.GET("/admin/", h.Home)
	r.GET("/admin/users/:id", h.GetUserById)
	r.POST("/admin/users", h.AddUser)
	r.PATCH("/admin/users/:id", h.UpdateUser)
	r.DELETE("/admin/users/:id", h.DeleteUser)
}
