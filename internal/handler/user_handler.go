package handler

import (
	"backend/internal/models"
	"context"
	"fmt"
	"net/http"
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

	conn, err := pgx.Connect(context.Background(), connConfig.ConnString())

	if err != nil {
		fmt.Println("err euy")
		fmt.Println(err)
		ctx.JSON(http.StatusBadGateway, test{
			Success: false,
			Message: "Something Gone Wrong",
		})
		return
	}

	rows, err := conn.Query(context.Background(), `
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

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.User])

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
