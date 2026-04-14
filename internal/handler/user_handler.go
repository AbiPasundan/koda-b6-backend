package handler

import (
	"backend/internal/helper"
	"backend/internal/models"
	"backend/internal/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: service,
	}
}

// Home godoc
//
//	@Summary		Get All Users
//	@Description	Retrieve all users from the database
//	@Tags			users
//	@Produce		json
//	@Success		200	{object}	models.Response
//	@Failure		500	{object}	models.Response
//	@Router			/users [get]
func (h *UserHandler) Home(ctx *gin.Context) {
	users, err := h.UserService.GetUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed get users",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Data User",
		Results: users,
	})
}

// SearchUser godoc
//
//	@Summary		Get user by ID
//	@Description	Retrieve a single user by its ID parameter
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	models.Response
//	@Failure		400	{object}	models.Response
//	@Failure		404	{object}	models.Response
//	@Router			/users/{id} [get]
func (h *UserHandler) GetUserById(ctx *gin.Context) {
	id, ok := helper.GetID(ctx)
	if !ok {
		return
	}
	user, err := h.UserService.GetUserById(id)
	if helper.NotFoundError(ctx, err) {
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User found:)))",
		Results: user,
	})
}

// Add User godoc
//
//	@Summary		AddUser Post
//	@Description	AddUser Process
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Users
//	@Failure		400	{object}	models.Users
//	@Router			/users [post]
func (h *UserHandler) AddUser(ctx *gin.Context) {
	var newUser models.User

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Something Went Wrong" + err.Error(),
			Results: nil,
		})
		return
	}
	createUser, err := h.UserService.AddUser(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Internal Server Error" + err.Error(),
			Results: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Success: false,
		Message: "Internal Server Error",
		Results: createUser,
	})
}

// DeleteUser godoc
//
//	@Summary		Delete user
//	@Description	Delete user by ID
//	@Tags			users
//	@Produce		json
//	@Param			id	path		int	true	"User ID"	Format(int64)
//	@Success		200	{object}	models.Response
//	@Failure		400	{object}	models.Response
//	@Failure		404	{object}	models.Response
//	@Router			/users/{id} [delete]
func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	id, ok := helper.GetID(ctx)
	if !ok {
		return
	}
	err := h.UserService.DeleteUserById(id)
	if helper.NotFoundError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: fmt.Sprintf("Success delete user with id: %d", id),
		Results: nil,
	})

}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	id, ok := helper.GetID(ctx)
	if !ok {
		return
	}

	var updateUser models.User
	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Something wrong" + err.Error(),
			Results: nil,
		})
		return
	}

	createUser, err := h.UserService.AddUser(updateUser)
	if helper.NotFoundError(ctx, err) {
		return
	}

	h.UserService.UpdateUserById(id, createUser)

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Succes Updated User",
		Results: createUser,
	})
}
