package router

import (
	"github.com/labstack/echo/v4"
	"github.com/trailstem/go-nextjs-rest-api/controller"
)

type Router struct {
	e *echo.Echo
}

func NewRouter() *Router {
	return &Router{
		e: echo.New(),
	}
}

func (r *Router) SetupRoutesAndStart(uc controller.IUserController) {

	r.e.POST("/signup", uc.SignUp)
	r.e.POST("/login", uc.LogIn)
	r.e.POST("/logout", uc.LogOut)
	// サーバー起動
	r.e.Logger.Fatal(r.e.Start(":8080"))
}
