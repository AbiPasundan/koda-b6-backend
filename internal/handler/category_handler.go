package handler

import (
	"backend/internal/helper"
	"backend/internal/models"
	"backend/internal/service"
	"fmt"

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
	statusInternal := helper.InternalServerError(ctx, "Internal Server Error", category, err)
	if statusInternal {
		return
	}

	helper.ResponseOk(ctx, "Success get All Data Category", &category)
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

	helper.ResponseOk(ctx, "Success get Data Category", &category)
}

func (h *CategoryHandler) AddCategory(ctx *gin.Context) {
	var newCategory models.Category

	err := ctx.ShouldBindJSON(&newCategory)
	badReq := helper.BadRequest(ctx, "Invalid request body", nil, err)
	if badReq {
		return
	}
	createCategory, err := h.CategoryService.AddCategory(newCategory)
	serverInternal := helper.InternalServerError(ctx, "Internal Server Error", createCategory, err)
	if serverInternal {
		return
	}
	helper.ResponseOk(ctx, "Success Add Category", createCategory)
}

func (h *CategoryHandler) UpdateCategory(ctx *gin.Context) {
	var Category models.Category

	id, ok := helper.GetID(ctx)
	if !ok {
		return
	}

	badRed := ctx.ShouldBindJSON(&Category)
	if helper.BadRequest(ctx, "Invalid request body", nil, badRed) {
		return
	}

	updatedCategory, err := h.CategoryService.UpdateCategoryById(id, Category)
	if helper.NotFoundError(ctx, err) {
		return
	}

	helper.ResponseOk(ctx, "Successfully Updated Category", updatedCategory)
}

func (h *CategoryHandler) DeleteCategory(ctx *gin.Context) {
	id, ok := helper.GetID(ctx)
	if !ok {
		return
	}

	err := h.CategoryService.DeleteCategoryById(id)
	if helper.NotFoundError(ctx, err) {
		return
	}

	helper.ResponseOk(ctx, fmt.Sprintf("Success delete category with id: %d", id), nil)
}
