package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// func CORSMiddleware() gin.HandlerFunc {
// 	godotenv.Load()
// 	return func(ctx *gin.Context) {
// 		ctx.Header("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
// 		ctx.Header("Access-Controll-Allow-Headers", "Content-Type")
// 		if ctx.Request.Method == "OPTIONS" {
// 			ctx.Data(http.StatusOK, "", []byte(""))
// 		} else {
// 			ctx.Next()
// 		}
// 	}
// }

func CORSMiddleware() gin.HandlerFunc {
	godotenv.Load()
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", fmt.Sprintf(os.Getenv("FRONTEND_URL"), "http://localhost:5173"))
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		ctx.Header("Access-Controll-Allow-Headers", "content-type,authorization")
		if ctx.Request.Method == "OPTIONS" {
			ctx.Data(http.StatusOK, "", []byte(""))
		} else {
			ctx.Next()
		}
	}
}
