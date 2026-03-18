package routes

import (
	"backend/internal/handler"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine, h *handler.CategoryHandler) {
	admin := r.Group("/admin")
	admin.Use(middleware.JWTMiddleware())
	{
		admin.GET("/category", h.Category)
		admin.GET("/category/:id", h.SearchCategoryById)
		admin.PATCH("/category/:id", h.UpdateCategory)
		admin.DELETE("/category/:id", h.DeleteCategory)
		admin.POST("/category", h.AddCategory)
	}
}
