package routes

import (
	"waysbean/handlers"
	"waysbean/pkg/middleware"
	"waysbean/pkg/mysql"
	"waysbean/repositories"

	"github.com/labstack/echo/v4"
)

func ProfileRoutes(e *echo.Group) {
	profileRepository := repositories.NewProfileRepository(mysql.DB)
	h := handlers.HandlerProfile(profileRepository)

	e.GET("/profiles", h.FindProfile)
	e.POST("/profile", h.CreateProfile)
	e.GET("/profile-user", middleware.Auth(h.GetProfileByUserID))
}
