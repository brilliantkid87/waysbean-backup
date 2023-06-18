package transactiondto

import (
	"time"
	"waysbean/models"
)

type TransactionResponse struct {
	ID         int                   `json:"id"`
	User       models.UserResponse   `json:"user"`
	Name       string                `json:"name"  form:"name"`
	Email      string                `json:"email"  form:"email"`
	Phone      string                `json:"phone"  form:"phone"`
	Address    string                `json:"address"  form:"address"`
	Attachment string                `json:"attachment"  form:"attachment"`
	Status     string                `json:"status"  form:"attachment"`
	SubTotal   int                   `json:"sub_total"`
	TotalQty   int                   `json:"total_qty"`
	CreatedAt  time.Time             `json:"created_at"`
	UpdatedAt  time.Time             `json:"updated_at"`
	Cart       []models.CartResponse `json:"products"`
}
