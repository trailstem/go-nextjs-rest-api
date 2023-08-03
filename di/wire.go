//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/trailstem/go-nextjs-rest-api/controller"
	"github.com/trailstem/go-nextjs-rest-api/repository"
	"github.com/trailstem/go-nextjs-rest-api/usecase"
	"gorm.io/gorm"
)

func InitializeApp(db *gorm.DB) (controller.IUserController, error) {
	wire.Build(repository.NewUserRepository, usecase.NewUserUsecase, controller.NewUserController)
	return nil, nil
}
