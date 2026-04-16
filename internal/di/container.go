package container

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Container struct {
	Pool            *pgxpool.Pool
	rdb             *redis.Client
	UserHandler     *handler.UserHandler
	ProductHandler  *handler.ProductHandler
	CategoryHandler *handler.CategoryHandler
	AuthHandler     *handler.AuthHandler
	AddToCart       *handler.ProductCartHandler
	ProfileHandler  *handler.ProfileHandler
	OrderHandler    *handler.OrderHandler
}

type configRedis struct {
	rdbhost     string
	rdbport     string
	rdbpassword string
}

func BuildContainer() *Container {
	redisAddr := os.Getenv("REDIS_HOST")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisPort := os.Getenv("REDIS_PORT")

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisAddr, redisPort), // (redisAddr),
		Password: redisPassword,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Redis tidak merespon: %v", err)
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

	userRepo := repository.NewUserRepository(pool, rdb)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	productRepo := repository.NewProductRepository(pool, rdb)
	productService := service.NewProductService(productRepo, rdb)
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

	profileRepo := repository.NewUserRepository(pool, rdb)
	profileService := service.NewProfileService(profileRepo)
	profileHandler := handler.NewProfileHandler(profileService)

	orderRepo := repository.NewOrderRepository(pool)
	OrderService := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(OrderService)

	return &Container{
		Pool:            pool,
		rdb:             rdb,
		UserHandler:     userHandler,
		ProductHandler:  productHandler,
		CategoryHandler: categoryHandler,
		AuthHandler:     authHandler,
		AddToCart:       addToCartHandler,
		ProfileHandler:  profileHandler,
		OrderHandler:    orderHandler,
	}
}
