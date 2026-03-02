package main

import (
	"backend/internal/handler"
	"backend/internal/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	r.GET("/", handler.Home)
	r.GET("/users/:id", handler.SearchUser)

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
