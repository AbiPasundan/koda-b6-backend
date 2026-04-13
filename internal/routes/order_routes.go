package routes

import (
	"backend/internal/handler"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine, h *handler.OrderHandler) {
	user := r.Group("")
	user.Use(middleware.JWTMiddleware(), middleware.RoleMiddleware("user", "admin"))
	user.GET("/orders", h.GetOrder)
	// user.GET("/historyorder", h.HistoryOrder)
	// user.GET("/historyorder/:id", h.GetOrderById)
	// user.POST("/checkout", h.AddOrder)
	// user.POST("/detailproduct/addcart", h.AddCart)
	// user.DELETE("/detailproduct/deletecart", h.DeleteCart)
}
