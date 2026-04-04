package handler

import (
	"backend/internal/helper"
	"backend/internal/models"
	"backend/internal/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
	authHeader := ctx.Request.Header.Get("Authorization")

	tokenString, found := strings.CutPrefix(authHeader, "Bearer ")
	if !found || tokenString == "" {
		ctx.JSON(401, gin.H{"error": "Missing or invalid token"})
		return
	}

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte("SECRET_KEY"), nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(401, gin.H{"error": "Invalid token"})
		return
	}

	userID := int(claims["user_id"].(float64))

	cart, err := h.ProductCartService.GetOrder(userID)
	if helper.NotFoundError(ctx, err) {
		return
	}

	helper.ResponseOk(ctx, "Success getting Order data", &cart)
}

func (h *ProductCartHandler) AddOrder(ctx *gin.Context) {
	test := ctx.Request.Context()

	userIDValue, exists := ctx.Get("user_id")
	if !exists {
		helper.CustomeError(ctx, http.StatusUnauthorized, "Unauthorized", nil, nil)
		return
	}

	userID, ok := userIDValue.(int)
	if !ok {
		helper.BadRequest(ctx, "Invalid user ID ", nil, nil)
		return
	}

	orderID, err := h.ProductCartService.AddOrder(test, userID)
	if err != nil {
		helper.BadRequest(ctx, "Cart Masi Kosong ", nil, err)
		return
	}

	helper.ResponseOk(ctx, "Success", orderID)
}
