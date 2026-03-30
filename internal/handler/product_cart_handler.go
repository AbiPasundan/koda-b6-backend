package handler

import (
	"backend/internal/helper"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductCartHandler struct {
	ProductCartService *service.ProductCartService
}

func NewProductCartHandler(service *service.ProductCartService) *ProductCartHandler {
	return &ProductCartHandler{
		ProductCartService: service,
	}
}

type AddToCartRequest struct {
	ProductID   int    `json:"product_id" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
	ProductName string `json:"product_name" binding:"required"`
	BasePrice   int    `json:"base_price" binding:"required"`
	VariantName string `json:"variant_name"`
	SizeName    string `json:"size_name"`
}

func (h *ProductCartHandler) AddToCart(ctx *gin.Context) {
	// var CartItem models.CartItem
	// // var cart
	// if err := ctx.ShouldBindJSON(&CartItem); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, models.Response{
	// 		Success: false,
	// 		Message: "Something Went Wrong" + err.Error(),
	// 		Results: nil,
	// 	})
	// 	helper.CustomeError(ctx, http.StatusInternalServerError, "Successfully added item to cart", err.Error(), err)
	// 	return
	// }

	// err := h.ProductCartService.AddCart(ctx, )
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, models.Response{
	// 		Success: false,
	// 		Message: "Internal Server Error" + err.Error(),
	// 		Results: nil,
	// 	})
	// 	return
	// }
	helper.ResponseOk(ctx, "Successfully added item to cart", nil)
}
