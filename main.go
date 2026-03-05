package main

import (
	_ "backend/docs"
	"backend/internal/handler"
	"backend/internal/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Coffee Shop API
// @version 1.0.0
// @description coffe shop project
// @termsOfService http://swagger.io/terms/

// @host localhost:8888
// @BasePath /api/v1
// @securityDefinitions.basic BasicAuths

func main() {

	godotenv.Load()

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	r.GET("/", handler.Home)
	r.GET("/users/:id", handler.SearchUser)
	r.DELETE("/users/:id", handler.DeleteUser)
	r.POST("/users", handler.AddUser)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
