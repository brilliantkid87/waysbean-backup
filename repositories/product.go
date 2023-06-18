package repositories

import (
	"waysbean/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProduct() ([]models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	GetProduct(id int) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	DeleteProduct(product models.Product, id int) (models.Product, error)
}

func NewProductRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// method array
func (r *repository) FindProduct() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error

	return products, err
}

func (r *repository) GetProduct(id int) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error

	return product, err
}

func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error

	return product, err
}

func (r *repository) UpdateProduct(product models.Product) (models.Product, error) {
	err := r.db.Save(&product).Error
	return product, err
}

func (r *repository) DeleteProduct(product models.Product, id int) (models.Product, error) {
	err := r.db.Delete(&product, id).Error
	return product, err
}
