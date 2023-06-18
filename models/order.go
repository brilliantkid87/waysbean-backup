package models

type Order struct {
	ID            int `json:"id"`
	ProductID     ProductResponse
	UserID        UserResponse
	TransactionID TransactionResponse
}
