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
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
