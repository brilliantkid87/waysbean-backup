package routes

import (
	"waysbean/handlers"
	"waysbean/pkg/middleware"
	"waysbean/pkg/mysql"
	"waysbean/repositories"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	productRepository := repositories.NewProductRepository(mysql.DB)
	h := handlers.HandlerProduct(productRepository)

	e.GET("/products", h.FindProduct)
	e.POST("/product", middleware.UploadFile(h.CreateProduct))
	e.GET("/product/:id", h.GetProduct)
}
