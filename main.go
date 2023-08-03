package main

import (
	"github.com/trailstem/go-nextjs-rest-api/controller"
	"github.com/trailstem/go-nextjs-rest-api/db"
	"github.com/trailstem/go-nextjs-rest-api/repository"
	"github.com/trailstem/go-nextjs-rest-api/router"
	"github.com/trailstem/go-nextjs-rest-api/usecase"
)

func main() {
	db := db.NewDB()
	userController := controller.NewUserController(usecase.NewUserUsecase(repository.NewUserRepository(db)))
	e := router.NewRouter(userController)
	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
