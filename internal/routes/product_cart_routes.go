package routes

import (
	"backend/internal/handler"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ProductCartRoutes(r *gin.Engine, h *handler.ProductCartHandler) {
	user := r.Group("")
	user.Use(middleware.JWTMiddleware(), middleware.RoleMiddleware("user", "admin"))
	user.POST("/detailproduct/addcart", h.AddCart)

	user.GET("/detailproduct/addcart/:id", h.GetCart)
	user.GET("/historyorder", h.HistoryOrder)
	user.POST("/checkout", h.AddOrder)
}
