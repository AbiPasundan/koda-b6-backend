package handler

import (
	"backend/internal/models"
	"backend/internal/service"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: service,
	}
}

func (h *UserHandler) Home(ctx *gin.Context) {

	godotenv.Load()
	conn, err := pgx.Connect(context.Background(), "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Database connection error",
		})
		return
	}
	defer conn.Close(context.Background())

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
	i := ctx.Param("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid Id" + err.Error(),
			Results: nil,
		})
		return
	}
	user, err := h.UserService.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "404 Not Found " + err.Error(),
			Results: nil,
		})
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
	i := ctx.Param("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid Id" + err.Error(),
			Results: nil,
		})
		return
	}
	h.UserService.DeleteUserById(id)
	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Succes delete user with id : ",
		Results: "result",
	})
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	var updateUser models.User
	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Hmmm eror naon ieu nya???" + err.Error(),
			Results: nil,
		})
	}

	i := ctx.Param("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid Id" + err.Error(),
			Results: nil,
		})
		return
	}
	createUser, err := h.UserService.AddUser(updateUser)
	h.UserService.UpdateUserById(id, createUser)
	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Succes Updated User",
		Results: createUser,
	})
}
