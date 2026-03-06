package routes

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, h *handler.UserHandler) {
	r.GET("/", h.Home)
	r.GET("/users/:id", handler.SearchUser)
	r.DELETE("/users/:id", handler.DeleteUser)
	r.POST("/users", handler.AddUser)
}
