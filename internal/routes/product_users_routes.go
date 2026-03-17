package routes

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func ProductUserRoutes(r *gin.Engine, h *handler.ProductHandler) {
	r.GET("/products/home", h.ProductHome)
}

// admin.GET("/products/home", h.ProductHome)
