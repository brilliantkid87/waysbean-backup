package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	UserRoutes(e)
	TransactionRoutes(e)
	ProductRoutes(e)
	CartRoutes(e)
	ProfileRoutes(e)
	Authroutes(e)
}
