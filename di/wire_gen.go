// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/trailstem/go-nextjs-rest-api/controller"
	"github.com/trailstem/go-nextjs-rest-api/repository"
	"github.com/trailstem/go-nextjs-rest-api/usecase"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeApp(db *gorm.DB) (*AppControllers, error) {
	iUserRepository := repository.NewUserRepository(db)
	iUserUsecase := usecase.NewUserUsecase(iUserRepository)
	iUserController := controller.NewUserController(iUserUsecase)
	iTaskRepository := repository.NewTaskRepository(db)
	iTaskUsecase := usecase.NewTaskUsecase(iTaskRepository)
	iTaskController := controller.NewTaskController(iTaskUsecase)
	appControllers := NewAppControllers(iUserController, iTaskController)
	return appControllers, nil
}

// wire.go:

// uc : userController
// tc : taskController
type AppControllers struct {
	UserController controller.IUserController
	TaskController controller.ITaskController
}

func NewAppControllers(uc controller.IUserController, tc controller.ITaskController) *AppControllers {
	return &AppControllers{
		UserController: uc,
		TaskController: tc,
	}
}
