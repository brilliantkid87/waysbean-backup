package routes

import (
	"waysbean/handlers"
	"waysbean/pkg/middleware"
	"waysbean/pkg/mysql"
	"waysbean/repositories"

	"github.com/labstack/echo/v4"
)

func Authroutes(e *echo.Group) {
	authRepository := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authRepository)

	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.GET("/check-auth", middleware.Auth(h.CheckAuth))
}
