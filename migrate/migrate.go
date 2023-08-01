package main

import (
	"fmt"

	"github.com/trailstem/go-nextjs-rest-api/db"
	"github.com/trailstem/go-nextjs-rest-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Succesfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
