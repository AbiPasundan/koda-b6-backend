package handler

import (
	"backend/internal/models"
	"backend/internal/service"
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
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Something went wrong please try again : " + err.Error(),
			Results: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Success get Data Product",
		Results: &product,
	})
}

func (h *ProductHandler) SearchProductById(ctx *gin.Context) {

	i := ctx.Param("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid ID : " + err.Error(),
			Results: nil,
		})
		return
	}

	var product models.Product
	product, err = h.ProductService.GetProductById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Something Gone Wrong : " + err.Error(),
			Results: nil,
		})
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

	i := ctx.Param("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid id: " + err.Error(),
			Results: nil,
		})
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
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to update product: " + err.Error(),
			Results: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Successfully Updated Product",
		Results: updatedProduct,
	})
}

func (h *ProductHandler) DeleteProduct(ctx *gin.Context) {
	i := ctx.Param("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid id: " + err.Error(),
			Results: nil,
		})
		return
	}
	h.ProductService.DeleteProductById(id)

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Successfully Delete Product",
	})

}
