package main

import (
	"log"

	"github.com/trailstem/go-nextjs-rest-api/db"
	"github.com/trailstem/go-nextjs-rest-api/di"
	"github.com/trailstem/go-nextjs-rest-api/router"
)

func main() {
	db := db.NewDB()
	//以下wireを使用しない場合のコンストラクタインジェクション
	// userController := controller.NewUserController(usecase.NewUserUsecase(repository.NewUserRepository(db)))

	//DIツールwireを使用してDI実現
	userController, err := di.InitializeApp(db)
	if err != nil {
		log.Fatalln(err)
	}
	e := router.NewRouter(userController)
	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
