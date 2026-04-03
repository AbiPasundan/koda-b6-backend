package handler

import (
	"backend/internal/helper"
	"backend/internal/models"
	"backend/internal/service"
	"net/http"

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

func (h *ProductCartHandler) AddCart(c *gin.Context) {
	var req models.AddCartRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Format data tidak valid atau ada data yang kurang",
			"error":   err.Error(),
		})
		return
	}

	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	req.UserID = userID.(int)

	err := h.ProductCartService.AddCart(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Gagal memproses keranjang belanja",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Produk berhasil ditambahkan ke keranjang",
	})
}

func (h *ProductCartHandler) GetCart(ctx *gin.Context) {
	id, ok := helper.GetID(ctx)
	if !ok {
		return
	}

	cart, err := h.ProductCartService.GetCart(id)
	if helper.NotFoundError(ctx, err) {
		return
	}

	helper.ResponseOk(ctx, "Success getting Cart data", &cart)
}
func (h *ProductCartHandler) HistoryOrder(ctx *gin.Context) {
	// id, ok := helper.GetID(ctx)
	// if !ok {
	// 	return
	// }

	cart, err := h.ProductCartService.GetOrder()
	if helper.NotFoundError(ctx, err) {
		return
	}

	helper.ResponseOk(ctx, "Success getting Order data", &cart)
}
