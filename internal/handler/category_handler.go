package handler

import (
	"backend/internal/helper"
	"backend/internal/models"
	"backend/internal/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	CategoryService *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		CategoryService: service,
	}
}

// category godoc
//
//	@Summary		Get All Category
//	@Description	Retrieve all category from the database
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Response
//	@Failure		500	{object}	models.Response
//	@Router			/ [get]
func (h *CategoryHandler) Category(ctx *gin.Context) {
	category, err := h.CategoryService.GetCategory()
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
		Message: "List of Category",
		Results: &category,
	})
}

func (h *CategoryHandler) SearchCategoryById(ctx *gin.Context) {
	id, ok := helper.GetID(ctx)
	if !ok {
		return
	}

	var category models.Category
	category, err := h.CategoryService.GetCategoryById(id)
	if helper.NotFoundError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User Found:)))",
		Results: &category,
	})
}

func (h *CategoryHandler) AddCategory(ctx *gin.Context) {
	var newCategory models.Category

	err := ctx.ShouldBindJSON(&newCategory)
	badReq := helper.BadRequest(ctx, "Invalid request body", nil, err)
	if badReq {
		return
	}
	// if err := ctx.ShouldBindJSON(&newCategory); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, models.Response{
	// 		Success: false,
	// 		Message: "Something Went Wrong " + err.Error(),
	// 		Results: nil,
	// 	})
	// 	return
	// }
	createCategory, err := h.CategoryService.AddCategory(newCategory)
	serverInternal := helper.InternalServerError(ctx, "Internal Server Error", createCategory, err)
	if serverInternal {
		return
	}
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, models.Response{
	// 		Success: false,
	// 		Message: "Internal Server Error" + err.Error(),
	// 		Results: nil,
	// 	})
	// 	return
	// }
	helper.ResponseOk(ctx, "Success Add Category", createCategory)
	// ctx.JSON(http.StatusOK, models.Response{
	// 	Success: true,
	// 	Message: "Success Add Category",
	// 	Results: createCategory,
	// })
}

func (h *CategoryHandler) UpdateCategory(ctx *gin.Context) {
	var Category models.Category

	id, ok := helper.GetID(ctx)
	if !ok {
		return
	}

	if err := ctx.ShouldBindJSON(&Category); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid request body: " + err.Error(),
			Results: nil,
		})
		return
	}
	updatedCategory, err := h.CategoryService.UpdateCategoryById(id, Category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to update Category: " + err.Error(),
			Results: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Successfully Updated Category",
		Results: updatedCategory,
	})
}

func (h *CategoryHandler) DeleteCategory(ctx *gin.Context) {
	id, ok := helper.GetID(ctx)
	if !ok {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid ID format",
		})
		return
	}

	err := h.CategoryService.DeleteCategoryById(id)
	if helper.NotFoundError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: fmt.Sprintf("Success delete category with id: %d", id),
		Results: nil,
	})
}
