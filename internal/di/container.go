package container

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
)

type Container struct {
	UserHandler *handler.UserHandler
}

type ProductContainer struct {
	ProductHandler *handler.ProductHandler
}

func BuildContainer() *Container {

	userRepo := repository.NewUserRepository()

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	return &Container{
		UserHandler: userHandler,
	}
}

func ProductsContainer() *ProductContainer {

	productRepo := repository.NewProductRepository()

	productService := service.NewProductService(productRepo)

	productHandler := handler.NewProductHandler(productService)

	return &ProductContainer{
		ProductHandler: productHandler,
	}
}
