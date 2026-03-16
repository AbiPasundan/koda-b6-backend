package main

import (
	_ "backend/cmd/docs"
	container "backend/internal/di"
	"backend/internal/middleware"
	"backend/internal/routes"
	"fmt"
	"os"

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
	// productContainer := container.ProductsContainer()
	// forgotPasswordContainer := container.ForgotPasswordContainer()

	routes.UserRoutes(r, userContainer.UserHandler)
	routes.ProductRoutes(r, userContainer.ProductHandler)
	routes.AuthRoutes(r, userContainer.ForgotPasswordHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))

	// r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}

// migrate create -ext sql -dir migrations -seq i init_db
// with time format
// migrate create -ext sql -dir migrations -format timestamp test_with_time_now

// migrate -source file://./migrations -database postgres://postgres:1@localhost:5432/yuuke?sslmode=disable up

// goseeder seed -s file://./seeder -d postgres://postgres:1@localhost:5432/postgres?sslmode=disable
