package main

import (
	_ "backend/docs"
	container "backend/internal/di"
	"backend/internal/handler"
	"backend/internal/middleware"
	"backend/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Coffee Shop API
// @version 1.0.0
// @description coffe shop project
// @termsOfService http://swagger.io/terms/

// @host localhost:PORT
// @BasePath /api/v1
// @securityDefinitions.basic BasicAuths

func main() {

	godotenv.Load()

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	c := container.BuildContainer()

	routes.UserRoutes(r, c.UserHandler)

	// r.GET("/", handler.Home)
	r.GET("/users/:id", handler.SearchUser)
	r.DELETE("/users/:id", handler.DeleteUser)
	r.POST("/users", handler.AddUser)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run("localhost:8888")

	// r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
