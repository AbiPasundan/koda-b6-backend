package routes

import (
	"backend/internal/handler"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ProductCartRoutes(r *gin.Engine, h *handler.ProductCartHandler) {
	user := r.Group("/user")
	user.Use(middleware.JWTMiddleware(), middleware.RoleMiddleware("user", "admin"))
	user.POST("/detailproduct/addcart", h.AddCart)
}

// {
//	"user_id": 5,
//	"product_id": 7,
//	"quantity": 3,
//	"product_name": "euy",
//	"base_price": 1000,
//	"variant_name": "no variant",
//	"size_name": "no size"
// }
