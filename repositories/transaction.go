package repositories

import (
	"waysbean/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransaction() ([]models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetTransaction(id int) (models.Transaction, error)
	GetTransactionByUser(ID int) ([]models.Transaction, error)
	GetTransactionByCartID(ID int) ([]models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

// method array
func (r *repository) FindTransaction() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("User").Preload("Cart.Product").Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").First(&transaction, ID).Error

	return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("User").Preload("Cart.Product").Create(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransactionByUser(ID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Where("user_id =?", ID).Preload("User").Preload("Cart.Product").Find(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransactionByCartID(ID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Where("cart_id = ?", ID).Preload("User").Find(&transactions).Error

	return transactions, err
}
