package routes

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, h *handler.UserHandler) {

	r.GET("/", h.Home)

}
