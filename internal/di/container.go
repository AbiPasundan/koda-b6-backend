package container

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type Container struct {
	UserHandler           *handler.UserHandler
	ProductHandler        *handler.ProductHandler
	CategoryHandler       *handler.CategoryHandler
	AuthHandler           *handler.AuthHandler
	ForgotPasswordHandler *handler.ForgotPasswordHandler
}

func BuildContainer() *Container {
	godotenv.Load()
	connConfig, err := pgx.ParseConfig("")
	if err != nil {
		log.Fatal(err.Error())
	}

	conn, err := pgx.Connect(context.Background(), connConfig.ConnString())
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepo := repository.NewUserRepository(conn)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	productRepo := repository.NewProductRepository(conn)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	categoryRepo := repository.NewCategoryRepository(conn)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	authRepo := repository.NewAuthRepository(conn)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	return &Container{
		UserHandler:     userHandler,
		ProductHandler:  productHandler,
		CategoryHandler: categoryHandler,
		AuthHandler:     authHandler,
	}
}

// type Container struct {
// 	UserHandler *handler.UserHandler
// }

// type ProductContainer struct {
// 	ProductHandler *handler.ProductHandler
// }

// type ForgotPassword struct {
// 	ForgotPasswordHandler *handler.ForgotPasswordHandler
// }

// func ForgotPasswordContainer() *ForgotPassword {
// 	godotenv.Load()
// 	connConfig, err := pgx.ParseConfig("")

// 	if err != nil {
// 		return nil
// 	}

// 	conn, err := pgx.Connect(context.Background(), connConfig.ConnString())
// 	if err != nil {
// 		return nil
// 	}

// 	forgotPasswordRepo := repository.NewForgotPasswordRepository(conn)
// 	userRepo := repository.NewUserRepository(conn)
// 	forgotPasswordService := service.NewForgotPasswordService(forgotPasswordRepo, userRepo)
// 	forgotPasswordHandler := handler.NewForgotPasswordHandler(forgotPasswordService)
// 	return &ForgotPassword{
// 		ForgotPasswordHandler: forgotPasswordHandler,
// 	}
// }

// func BuildContainer() *Container {
// 	godotenv.Load()
// 	connConfig, err := pgx.ParseConfig("")
// 	if err != nil {
// 		return nil
// 	}

// 	conn, err := pgx.Connect(context.Background(), connConfig.ConnString())
// 	if err != nil {
// 		return nil
// 	}

// 	userRepo := repository.NewUserRepository(conn)

// 	userService := service.NewUserService(userRepo)

// 	userHandler := handler.NewUserHandler(userService)

// 	return &Container{
// 		UserHandler: userHandler,
// 	}
// }

// func ProductsContainer() *ProductContainer {
// 	godotenv.Load()
// 	connConfig, err := pgx.ParseConfig("")
// 	if err != nil {
// 		return nil
// 	}

// 	conn, err := pgx.Connect(context.Background(), connConfig.ConnString())
// 	if err != nil {
// 		return nil
// 	}

// 	productRepo := repository.NewProductRepository(conn)

// 	productService := service.NewProductService(productRepo)

// 	productHandler := handler.NewProductHandler(productService)

// 	return &ProductContainer{
// 		ProductHandler: productHandler,
// 	}
// }
