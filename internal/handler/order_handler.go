package handler

import (
	"backend/internal/helper"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderService *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{
		OrderService: service,
	}
}

func (h *OrderHandler) GetOrder(ctx *gin.Context) {
	product, err := h.OrderService.GetOrder()
	if helper.InternalServerError(ctx, "Internal Server Error", product, err) {
		return
	}

	helper.ResponseOk(ctx, "Success get Data Product", &product)
}
