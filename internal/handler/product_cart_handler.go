package handler

import (
	"backend/internal/helper"
	"backend/internal/models"
	"backend/internal/service"
	"fmt"
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

func (h *ProductCartHandler) AddCart(ctx *gin.Context) {
	var req models.AddCartRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.BadRequest(ctx, "Invalid request body", nil, err)
		return
	}

	userID, exists := ctx.Get("user_id")

	if !exists {
		helper.CustomeError(ctx, http.StatusUnauthorized, "Unauthorized", nil, nil)
		return
	}
	req.UserID = userID.(int)

	err := h.ProductCartService.AddCart(ctx.Request.Context(), req)
	if err != nil {
		helper.InternalServerError(ctx, "Unauthorized", err.Error(), err)
		return
	}

	helper.ResponseOk(ctx, "Success Delete Cart Item", nil)
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

func (h *ProductCartHandler) DeleteCart(ctx *gin.Context) {
	var id models.DeleteCartItem

	if err := ctx.ShouldBindJSON(&id); err != nil {
		helper.BadRequest(ctx, "Invalid request body", nil, err)
		return
	}

	err := h.ProductCartService.DeleteCartById(id.ProductID)
	if helper.NotFoundError(ctx, err) {
		return
	}

	helper.ResponseOk(ctx, fmt.Sprintf("Success delete cart with id: %d", id.ProductID), nil)
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
	fmt.Println(cart)

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

func (h *ProductCartHandler) GetOrderById(ctx *gin.Context) {
	i := ctx.Param("id")

	cart, err := h.ProductCartService.GetOrderById(i)
	if helper.NotFoundError(ctx, err) {
		return
	}

	helper.ResponseOk(ctx, "Success getting Order data", &cart)
}
