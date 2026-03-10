package handler

import (
	"backend/internal/models"
	"backend/internal/service"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var ListUser []models.User

var users []models.Users
var user []models.User
var rows pgx.Rows
var conn *pgx.Conn

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
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Something went wrong" + err.Error(),
			Results: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User found",
		Results: user,
	})
}

// func SearchUser(ctx *gin.Context) {
// 	i := ctx.Param("id")
// 	id, _ := strconv.Atoi(i)
// 	connConfig, err := pgx.ParseConfig("")
// 	conn, err := pgx.Connect(context.Background(), connConfig.ConnString())

// 	rows, err = conn.Query(context.Background(), `
// 			SELECT id, full_name, email, password, address, phone, pictures FROM users WHERE id = $1
// 		`, id)
// 	user, err = pgx.CollectRows(rows, pgx.RowToStructByName[models.User])

// 	// rows, err = conn.Query(context.Background(), `
// 	// 		SELECT id, full_name, email, address, phone FROM users WHERE id = $1
// 	// 	`, id)
// 	// users, err = pgx.CollectRows(rows, pgx.RowToStructByName[models.Users])

// 	if err != nil {
// 		fmt.Println("err take data")
// 		fmt.Println(err)
// 		ctx.JSON(http.StatusBadRequest, models.Response{
// 			Success: false,
// 			Message: "Data User Gangguan",
// 		})
// 		return
// 	}

// 	if len(user) <= 0 {
// 		ctx.JSON(http.StatusNotFound, models.Response{
// 			Success: false,
// 			Message: "Data User Not Found",
// 			Results: nil,
// 		})
// 		return
// 	} else {
// 		ctx.JSON(http.StatusOK, models.Response{
// 			Success: true,
// 			Message: "Data User",
// 			Results: user,
// 		})
// 	}
// }

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
func AddUser(ctx *gin.Context) {
	connConfig, err := pgx.ParseConfig("")
	conn, err := pgx.Connect(context.Background(), connConfig.ConnString())

	rows, err = conn.Query(context.Background(), `
		INSERT INTO
			users
			("id", "full_name", "email", "password", "address", "phone", "pictures")
			VALUES
			(DEFAULT, 'New User From API','newuser@mail.com','newuser#123','Maharaja, Depok','0811234455','images/Response/path/user.jpg');
	`)

	users, err = pgx.CollectRows(rows, pgx.RowToStructByName[models.Users])

	if err != nil {
		fmt.Println("err take data")
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Data User",
		})
		return
	}
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
func DeleteUser(ctx *gin.Context) {

	rawId := ctx.Param("id")

	id, err := strconv.Atoi(rawId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid ID",
			Results: nil,
		})
		return
	}

	connConfig, err := pgx.ParseConfig("")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Database config error",
			Results: nil,
		})
		return
	}

	conn, err := pgx.Connect(context.Background(), connConfig.ConnString())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Database connection error",
			Results: nil,
		})
		return
	}

	defer conn.Close(context.Background())

	cmdTag, err := conn.Exec(context.Background(),
		`DELETE FROM users WHERE id = $1`, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed delete user",
			Results: nil,
		})
		return
	}

	if cmdTag.RowsAffected() == 0 {
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "User Not Found",
			Results: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User Deleted Successfully",
		Results: nil,
	})

}
