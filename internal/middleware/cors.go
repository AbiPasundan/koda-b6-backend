package middleware

import (
	"net/http"

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

// func CORSMiddleware() gin.HandlerFunc {
// 	godotenv.Load()
// 	allowedOrigin := os.Getenv("FRONTEND_URL")
// 	if allowedOrigin == "" {
// 		allowedOrigin = "http://localhost:5173"
// 	}

// 	return func(ctx *gin.Context) {
// 		ctx.Header("Access-Control-Allow-Origin", allowedOrigin)
// 		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
// 		ctx.Header("Access-Control-Allow-Headers", "content-type,authorization")
// 		ctx.Header("Access-Control-Allow-Credentials", "true")

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
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "content-type,authorization")
		if ctx.Request.Method == "OPTIONS" {
			ctx.Data(http.StatusOK, "", []byte(""))
			return
		} else {
			ctx.Next()
		}
	}
}
