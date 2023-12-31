//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/trailstem/go-nextjs-rest-api/controller"
	"github.com/trailstem/go-nextjs-rest-api/repository"
	"github.com/trailstem/go-nextjs-rest-api/usecase"
	"github.com/trailstem/go-nextjs-rest-api/validator"
	"gorm.io/gorm"
)

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

func InitializeApp(db *gorm.DB) (*AppControllers, error) {
	wire.Build(repository.NewUserRepository, validator.NewUserValidator, usecase.NewUserUsecase, controller.NewUserController,
		repository.NewTaskRepository, validator.NewTaskValidator, usecase.NewTaskUsecase, controller.NewTaskController, NewAppControllers)

	return &AppControllers{}, nil
}
