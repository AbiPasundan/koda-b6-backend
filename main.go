package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type test struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

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
	r := gin.Default()
	r.Use(CORSMiddleware())
	godotenv.Load()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, test{
			Success: true,
			Message: "Data User",
		})
	})

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
