package handler

import (
	"backend/internal/models"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type test struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any
}

var ListUser []models.User
var Counter int64

var users []models.Users
var user []models.User
var rows pgx.Rows
var conn *pgx.Conn

func idCounter() int64 {
	return atomic.AddInt64(&Counter, 1)
}

func Home(ctx *gin.Context) {
	connConfig, err := pgx.ParseConfig("")

	if err != nil {
		fmt.Println("err euy")
		fmt.Println(err)
		ctx.JSON(http.StatusBadGateway, test{
			Success: false,
			Message: "Something Gone Wrong",
		})
		return
	}

	conn, err = pgx.Connect(context.Background(), connConfig.ConnString())

	if err != nil {
		fmt.Println("err euy")
		fmt.Println(err)
		ctx.JSON(http.StatusBadGateway, test{
			Success: false,
			Message: "Something Gone Wrong",
		})
		return
	}

	rows, err = conn.Query(context.Background(), `
			SELECT id, full_name, email, address, phone FROM users
		`)

	if err != nil {
		fmt.Println("err euy")
		fmt.Println(err)
		ctx.JSON(http.StatusBadGateway, test{
			Success: false,
			Message: "Something Gone Wrong",
		})
		return
	}

	users, err = pgx.CollectRows(rows, pgx.RowToStructByName[models.Users])

	if err != nil {
		fmt.Println("err take data")
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, test{
			Success: false,
			Message: "Data User",
		})
		return
	}

	ctx.JSON(http.StatusOK, test{
		Success: true,
		Message: "Data User",
		Results: users,
	})

}

func SearchUser(ctx *gin.Context) {
	i := ctx.Param("id")
	id, _ := strconv.Atoi(i)
	connConfig, err := pgx.ParseConfig("")
	conn, err := pgx.Connect(context.Background(), connConfig.ConnString())

	rows, err = conn.Query(context.Background(), `
			SELECT id, full_name, email, password, address, phone, pictures FROM users WHERE id = $1
		`, id)
	user, err = pgx.CollectRows(rows, pgx.RowToStructByName[models.User])

	// rows, err = conn.Query(context.Background(), `
	// 		SELECT id, full_name, email, address, phone FROM users WHERE id = $1
	// 	`, id)
	// users, err = pgx.CollectRows(rows, pgx.RowToStructByName[models.Users])

	if err != nil {
		fmt.Println("err take data")
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, test{
			Success: false,
			Message: "Data User Gangguan",
		})
		return
	}

	if len(user) <= 0 {
		ctx.JSON(http.StatusNotFound, test{
			Success: false,
			Message: "Data User Not Found",
			Results: nil,
		})
		return
	} else {
		ctx.JSON(http.StatusOK, test{
			Success: true,
			Message: "Data User",
			Results: user,
		})
	}
}

func AddUser(ctx *gin.Context) {
	connConfig, err := pgx.ParseConfig("")
	conn, err := pgx.Connect(context.Background(), connConfig.ConnString())

	rows, err = conn.Query(context.Background(), `
		INSERT INTO
			users
			("id", "full_name", "email", "password", "address", "phone", "pictures")
			VALUES
			(DEFAULT, 'New User From API','newuser@mail.com','newuser#123','Maharaja, Depok','0811234455','images/test/path/user.jpg');
	`)

	users, err = pgx.CollectRows(rows, pgx.RowToStructByName[models.Users])

	if err != nil {
		fmt.Println("err take data")
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, test{
			Success: false,
			Message: "Data User",
		})
		return
	}

}

func DeleteUser(ctx *gin.Context) {
	// there is still a bug in delete
	// if user delete user using not available id
	// it still show messege true
	rawId := ctx.Param("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, test{
			Success: false,
			Message: "Id tidak ditemukan",
		})
	}
	connConfig, err := pgx.ParseConfig("")
	conn, err := pgx.Connect(context.Background(), connConfig.ConnString())

	rows, err = conn.Query(context.Background(), `
		DELETE FROM users WHERE id = $1
	`, id)

	users, err = pgx.CollectRows(rows, pgx.RowToStructByName[models.Users])

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, test{
			Success: false,
			Message: "Id tidak ditemukan",
		})
		return
	}

	// for _, v := range users {
	// 	switch v.Id {
	// 	case id:
	// 		ctx.JSON(http.StatusOK, test{
	// 			Success: true,
	// 			Message: "Data User Deleted",
	// 			Results: nil,
	// 		})
	// 	default:
	// 		ctx.JSON(http.StatusNotFound, test{
	// 			Success: false,
	// 			Message: "Data User Not Found",
	// 			Results: nil,
	// 		})
	// 	}
	// }

	for _, v := range users {
		if id != int(v.Id) {
			ctx.JSON(http.StatusNotFound, test{
				Success: false,
				Message: "Data User Not Found",
				Results: nil,
			})
			return
		}

	}

	ctx.JSON(http.StatusOK, test{
		Success: true,
		Message: "Data User Deleted",
		Results: users,
	})

	// if len(users) <= 0 {
	// 	ctx.JSON(http.StatusNotFound, test{
	// 		Success: false,
	// 		Message: "Data User Not Found",
	// 		Results: nil,
	// 	})
	// 	return
	// } else {
	// 	ctx.JSON(http.StatusOK, test{
	// 		Success: true,
	// 		Message: "Data User Deleted",
	// 		Results: users,
	// 	})
	// }
}
