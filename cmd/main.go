package main

import (
	_ "backend/cmd/docs"
	container "backend/internal/di"
	"backend/internal/middleware"
	"backend/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Coffee Shop API
//	@version		1.0.0
//	@description	coffe shop project
//	@host			localhost:8889
//	@BasePath		/api/v1

func main() {

	godotenv.Load()

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	userContainer := container.BuildContainer()
	productContainer := container.ProductsContainer()

	routes.UserRoutes(r, userContainer.UserHandler)
	routes.ProductRoutes(r, productContainer.ProductHandler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run("localhost:8889")

	// r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
