package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		// c.Writer.Header().Set("Access-Control-Allow-Origin", "/")
// 		// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		// c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		// c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

// 		ctx.Header("Access-Controll-Allow-Origin", "localhost:8888")
// 		ctx.Header("Access-Controll-Allow-Headers", "content-type")

// 		if ctx.Request.Method == "OPTIONS" {
// 			ctx.AbortWithStatus(204)
// 			return
// 		}
// 		ctx.Next()
// 	}
// }

func CORSMiddleware() gin.HandlerFunc {
	godotenv.Load()
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "localhost:5173")
		// ctx.Header("Access-Control-Allow-Origin", "content-type")
		if ctx.Request.Method == "OPTIONS" {
			ctx.Data(http.StatusOK, "", []byte(""))
		} else {
			ctx.Next()
		}
	}
}

type test struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
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
