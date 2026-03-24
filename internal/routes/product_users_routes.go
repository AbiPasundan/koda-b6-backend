package routes

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func ProductUserRoutes(r *gin.Engine, h *handler.ProductHandler) {
	// in landing page
	r.GET("/products/home", h.ProductHome)
	r.GET("/products/reviews", h.ProductReview)
	// in browse product
	r.GET("/products", h.Product)
	r.GET("/browseproducts", h.BrowseProduct)
}
