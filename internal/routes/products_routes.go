package routes

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, h *handler.ProductHandler) {
	r.GET("/admin/products", h.Product)
	r.GET("/admin/products/:id", h.SearchProductById)
	r.PATCH("/admin/products/:id", h.UpdateProduct)
	r.POST("/admin/products", h.AddProduct)
}
