package routes

import (
	"waysbean/handlers"
	"waysbean/pkg/mysql"
	"waysbean/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	e.GET("/users", h.FindUsers)
	// e.GET("/user/:id", h.GetUser)
	e.POST("/user", h.CreateUser)
	// e.PATCH("/user/:id", middleware.Auth(h.UpdateUser))
	// e.DELETE("/user/:id", middleware.Auth(h.DeleteUser))
}
