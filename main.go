package main

import (
	"fmt"
	"os"
	"waysbean/database"
	"waysbean/pkg/mysql"
	"waysbean/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	mysql.DatabaseConnection()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))
	e.Static("/uploads", "./uploads")

	PORT := os.Getenv("PORT")

	fmt.Println("server running localhost:" + PORT)
	e.Logger.Fatal(e.Start(":" + PORT))
}
