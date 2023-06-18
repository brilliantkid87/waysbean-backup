package database

import (
	"fmt"
	"waysbean/models"
	"waysbean/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Product{},
		&models.Transaction{},
		&models.Cart{},
	)

	if err != nil {
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
