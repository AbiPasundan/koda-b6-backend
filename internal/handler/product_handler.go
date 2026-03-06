package handler

import (
	"backend/internal/models"
	"backend/internal/service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type ProductHandler struct {
	ProductService *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: service,
	}
}

func (h *ProductHandler) Product(ctx *gin.Context) {
	godotenv.Load()
	conn, err := pgx.Connect(context.Background(), "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Something went wrong please try again",
			Results: nil,
		})
	}
	defer conn.Close(context.Background())
	product, err := h.ProductService.GetProduct(conn)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Something went wrong please try again",
			Results: nil,
		})
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Succes get Data Product",
		Results: product,
	})

}
