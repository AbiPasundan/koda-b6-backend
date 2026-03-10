package routes

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, h *handler.ProductHandler) {
	r.GET("/products", h.Product)
	r.GET("/products/:id", h.SearchProductById)
	r.POST("/products", h.AddProduct)
}
