package container

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type Container struct {
	UserHandler *handler.UserHandler
}

type ProductContainer struct {
	ProductHandler *handler.ProductHandler
}

func BuildContainer() *Container {
	godotenv.Load()
	connConfig, err := pgx.ParseConfig("")
	if err != nil {
		return nil
	}

	conn, err := pgx.Connect(context.Background(), connConfig.ConnString())
	if err != nil {
		return nil
	}

	userRepo := repository.NewUserRepository(conn)

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	return &Container{
		UserHandler: userHandler,
	}
}

func ProductsContainer() *ProductContainer {
	godotenv.Load()
	connConfig, err := pgx.ParseConfig("")
	if err != nil {
		return nil
	}

	conn, err := pgx.Connect(context.Background(), connConfig.ConnString())
	if err != nil {
		return nil
	}

	productRepo := repository.NewProductRepository(conn)

	productService := service.NewProductService(productRepo)

	productHandler := handler.NewProductHandler(productService)

	return &ProductContainer{
		ProductHandler: productHandler,
	}
}
