package repositories

import (
	"waysbean/models"

	"gorm.io/gorm"
)

// kontrak
type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(id int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

// method array
func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Profile").Preload("Transaction.Cart").Find(&users).Error

	return users, err
}

func (r *repository) GetUser(id int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}
