package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type test struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any
}

type UserBeta struct {
	Id        int    `json:"id"`
	Full_Name string `json:"full_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Pictures  string `json:"pictures"`
}

type user struct {
	Full_Name string `json:"full_name"`
	Email     string `json:"email"`
}

// select users.id, users.full_name, users.email, users.password, users.address, users.phone, users.pictures from users;

func CORSMiddleware() gin.HandlerFunc {
	godotenv.Load()
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Controll-Allow-Headers", "Content-Type")
		if ctx.Request.Method == "OPTIONS" {
			ctx.Data(http.StatusOK, "", []byte(""))
		} else {
			ctx.Next()
		}
	}
}

func main() {

	godotenv.Load()

	connConfig, err := pgx.ParseConfig("")

	if err != nil {
		fmt.Println("err euy")
		fmt.Println(err)
		return
	}

	conn, err := pgx.Connect(context.Background(), connConfig.ConnString())

	if err != nil {
		fmt.Println("err euy")
		fmt.Println(err)
		return
	}

	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/", func(ctx *gin.Context) {
		// query 1 ` select users.id, users.full_name, users.email, users.password, users. address, users.phone, users.pictures from users; `
		rows, err := conn.Query(context.Background(), `
			SELECT full_name, email FROM users
		`)

		if err != nil {
			fmt.Println("err take data")
			fmt.Println(err)
			return
		}

		users, err := pgx.CollectRows(rows, pgx.RowToStructByName[user])

		if err != nil {
			fmt.Println("err take data")
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, test{
				Success: true,
				Message: "Data User",
			})
			return
		}

		ctx.JSON(http.StatusOK, test{
			Success: true,
			Message: "Data User",
			Results: users,
		})
	})

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
