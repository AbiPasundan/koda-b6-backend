package handler

import (
	"backend/internal/helper"
	"backend/internal/models"
	"backend/internal/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductService *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: service,
	}
}

// Product godoc
//
//	@Summary		Get All Product
//	@Description	Retrieve all products from the system
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Response
//	@Failure		500	{object}	models.Response
//	@Router			/ [get]
func (h *ProductHandler) Product(ctx *gin.Context) {
	product, err := h.ProductService.GetProduct()
	if helper.InternalServerError(ctx, "Internal Server Error", product, err) {
		return
	}

	helper.ResponseOk(ctx, "Success get Data Product", &product)
}

// BrowseProduct godoc
//
//	@Summary		Get Browse Product
//	@Description	Retrieve Browse Product from the system
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Response
//	@Failure		500	{object}	models.Response
//	@Router			/browseproduct [get]
func (h *ProductHandler) BrowseProduct(ctx *gin.Context) {
	product, err := h.ProductService.BrowseProducts(ctx)
	if helper.InternalServerError(ctx, "Internal Server Error", product, err) {
		return
	}

	helper.ResponseOk(ctx, "Success get browse Data Product", &product)
}

func (h *ProductHandler) DetailProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	product, err := h.ProductService.DetailProduct(ctx, id)
	if helper.InternalServerError(ctx, "Internal Server Error", product, err) {
		return
	}

	helper.ResponseOk(ctx, "Success get browse Data Product", &product)
}

func (h *ProductHandler) ProductHome(ctx *gin.Context) {
	product, err := h.ProductService.GetProductHome(ctx)
	if helper.InternalServerError(ctx, "Internal Server Error", product, err) {
		return
	}

	helper.ResponseOk(ctx, "Success get Data Product", &product)
}

func (h *ProductHandler) ProductReview(ctx *gin.Context) {
	product, err := h.ProductService.ProductReview(ctx)
	if helper.InternalServerError(ctx, "Internal Server Error", product, err) {
		return
	}

	helper.ResponseOk(ctx, "Success get Review Product", &product)
}

func (h *ProductHandler) SearchProductById(ctx *gin.Context) {

	id, ok := helper.GetID(ctx)
	if !ok {
		return
	}

	var product models.Product
	product, err := h.ProductService.GetProductById(id)
	if helper.NotFoundError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User Found:)))",
		Results: &product,
	})
}

func (h *ProductHandler) AddProduct(ctx *gin.Context) {
	var newProducts models.Product

	if err := ctx.ShouldBindJSON(&newProducts); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Something Went Wrong" + err.Error(),
			Results: nil,
		})
		return
	}
	createUser, err := h.ProductService.AddProduct(newProducts)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Internal Server Error" + err.Error(),
			Results: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Success Add Product",
		Results: createUser,
	})
}

func (h *ProductHandler) UpdateProduct(ctx *gin.Context) {
	var product models.Product

	id, ok := helper.GetID(ctx)
	if !ok {
		return
	}
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid request body: " + err.Error(),
			Results: nil,
		})
		return
	}
	updatedProduct, err := h.ProductService.UpdateProductById(id, product)
	if helper.NotFoundError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Successfully Updated Product",
		Results: updatedProduct,
	})
}

func (h *ProductHandler) DeleteProduct(ctx *gin.Context) {
	id, ok := helper.GetID(ctx)
	if !ok {
		return
	}

	err := h.ProductService.DeleteProductById(id)
	if helper.NotFoundError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: fmt.Sprintf("Success delete product with id: %d", id),
		Results: nil,
	})

}
