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

	defer userContainer.Pool.Close()

	routes.UserRoutes(r, userContainer.UserHandler)
	routes.ProductRoutes(r, userContainer.ProductHandler)
	routes.ProductUserRoutes(r, userContainer.ProductHandler)
	routes.AuthRoutes(r, userContainer.AuthHandler)
	// routes.AuthRoutes(r, userContainer.ForgotPasswordHandler)
	routes.CategoryRoutes(r, userContainer.CategoryHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// it should be change when it test in prod
	r.Run(":8089")
	// r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))

	// r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}

// migrate create -ext sql -dir migrations -seq i init_db
// with time format
// migrate create -ext sql -dir migrations -format timestamp test_with_time_now

// migrate -source file://./migrations -database postgres://postgres:1@localhost:5432/yuuke?sslmode=disable up

// goseeder seed -s file://./seeder -d postgres://postgres:1@localhost:5432/postgres?sslmode=disable

//  docker pull ghcr.io/abipasundan/koda-b6-react:latest
//  docker rm -f web-wildan
//  docker run -d --network pman_web -p  20601:80 --name web-wildan ghcr.io/abipasundan/koda-b6-react:latest

//  cd server-reza
//  docker compose pull
//  docker run --rm --network=server-reza_default migrate/migrate:latest -source github:// rezafauzan:${{ secrets.GITHUB_TOKEN }}@rezafauzan/koda-b6-backend/migrations -database postgresql://postgres:1@pg:5432/postgres?sslmode=disable up
//  docker compose up backend -d

//  .github

//  docker build -t ghcr.io/YOUR_GITHUB_USERNAME/IMAGE_NAME:latest .
//  echo YOUR_PAT | docker login ghcr.io -u YOUR_GITHUB_USERNAME --password-stdin

// docker run --rm --network=server-wildan_default migrate/migrate:latest -source github://abipasundan/koda-b6-backend/migrations -database postgresql://postgres:1@pg:5432/postgres?sslmode=disable version

// perbaikan =
