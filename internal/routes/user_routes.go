package routes

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, h *handler.UserHandler) {
	r.GET("/users", h.Home)
	r.GET("/users/:id", h.GetUserById)
	r.POST("/users", h.AddUser)
	r.PATCH("/users/:id", h.UpdateUser)
	r.DELETE("/users/:id", h.DeleteUser)
}
