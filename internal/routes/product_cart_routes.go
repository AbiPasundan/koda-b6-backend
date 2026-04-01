package routes

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func ProductCartRoutes(r *gin.Engine, h *handler.ProductCartHandler) {
	// in landing page
	r.POST("/detailproduct/addcart", h.AddCart)
}
