package container

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
)

type Container struct {
	UserHandler *handler.UserHandler
}

func BuildContainer() *Container {

	userRepo := repository.NewUserRepository()

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	return &Container{
		UserHandler: userHandler,
	}
}
