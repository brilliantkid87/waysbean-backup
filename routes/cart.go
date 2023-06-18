package routes

import (
	"waysbean/handlers"
	"waysbean/pkg/mysql"
	"waysbean/repositories"

	"github.com/labstack/echo/v4"
)

func CartRoutes(e *echo.Group) {
	cartRepository := repositories.NewCartRepository(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	e.GET("/carts", h.FindCart)
	e.POST("/cart", h.CreateCart)
	e.GET("/cart-user/:id", h.GetCardByUser)
}
