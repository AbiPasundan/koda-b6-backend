package container

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Container struct {
	Pool            *pgxpool.Pool
	UserHandler     *handler.UserHandler
	ProductHandler  *handler.ProductHandler
	CategoryHandler *handler.CategoryHandler
	AuthHandler     *handler.AuthHandler
	AddToCart       *handler.ProductCartHandler
	ProfileHandler  *handler.ProfileHandler
	OrderHandler    *handler.OrderHandler
}

func BuildContainer() *Container {
	// godotenv.Load()
	// connConfig, err := pgx.ParseConfig("")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// conn, err := pgx.Connect(context.Background(), connConfig.ConnString())
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// godotenv.Load()
	// dbURL := os.Getenv("DATABASE_URL")

	// pool, err := pgxpool.New(context.Background(), dbURL)
	// if err != nil {
	// 	panic(err)
	// }
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatalf("Gagal parse config: %v", err)
	}

	config.MaxConns = 10

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Gagal membuat connection pool: %v", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("Database tidak merespon: %v", err)
	}
	log.Println("Berhasil terhubung menggunakan connection pool!")

	userRepo := repository.NewUserRepository(pool)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	productRepo := repository.NewProductRepository(pool)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	categoryRepo := repository.NewCategoryRepository(pool)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	authRepo := repository.NewAuthRepository(pool)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	addToCartRepo := repository.NewProductCartRepository(pool)
	addToCartService := service.NewProductCartService(addToCartRepo)
	addToCartHandler := handler.NewProductCartHandler(addToCartService)

	profileRepo := repository.NewUserRepository(pool)
	profileService := service.NewProfileService(profileRepo)
	profileHandler := handler.NewProfileHandler(profileService)

	orderRepo := repository.NewOrderRepository(pool)
	OrderService := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(OrderService)

	return &Container{
		UserHandler:     userHandler,
		ProductHandler:  productHandler,
		CategoryHandler: categoryHandler,
		AuthHandler:     authHandler,
		AddToCart:       addToCartHandler,
		ProfileHandler:  profileHandler,
		OrderHandler:    orderHandler,
	}
}
