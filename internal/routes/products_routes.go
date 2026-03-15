package routes

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, h *handler.ProductHandler) {
	admin := r.Group("/admin")
	{
		admin.GET("/products", h.Product)
		admin.GET("/products/:id", h.SearchProductById)
		admin.PATCH("/products/:id", h.UpdateProduct)
		admin.DELETE("/products/:id", h.DeleteProduct)
		admin.POST("/products", h.AddProduct)
	}
}
