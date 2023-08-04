package router

import (
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/trailstem/go-nextjs-rest-api/di"
)

type Router struct {
	e *echo.Echo
}

func NewRouter() *Router {
	return &Router{
		e: echo.New(),
	}
}

func (r *Router) SetupRoutesAndStart(ac *di.AppControllers) {

	// CORS
	r.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))

	//CSRF
	r.e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		// CookieSameSite: http.SameSiteDefaultMode, //postman動作確認用設定
		// CookieMaxAge: 60,
	}))

	r.e.POST("/signup", ac.UserController.SignUp)
	r.e.POST("/login", ac.UserController.LogIn)
	r.e.POST("/logout", ac.UserController.LogOut)
	r.e.GET("csrf", ac.UserController.CsrfToken)

	// ミドルウェア適用
	// SigningKey:JWT作成時と同じkey指定
	// TokenLookup: clientから送られてくるtoken格納場所を指定
	t := r.e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", ac.TaskController.GetAllTasks)
	t.GET("/:taskId", ac.TaskController.GetTaskById)
	t.POST("", ac.TaskController.CreateTask)
	t.PUT("/:taskId", ac.TaskController.UpdateTask)
	t.DELETE("/:taskId", ac.TaskController.DeleteTask)

	// サーバー起動
	r.e.Logger.Fatal(r.e.Start(":8080"))
}
