package routes

import (
	"backend/internal/handler"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ProductCartRoutes(r *gin.Engine, h *handler.ProductCartHandler) {
	r.GET("/detailproduct/addcart/:id", h.GetCart)
	user := r.Group("/user")
	user.Use(middleware.JWTMiddleware(), middleware.RoleMiddleware("user", "admin"))
	user.POST("/detailproduct/addcart", h.AddCart)
}
