package transactiondto

import "waysbean/models"

type CreateTransactionRequest struct {
	UserID   int                 `json:"user_id" form:"user_id"`
	User     models.UserResponse `json:"user"`
	Name     string              `json:"name" form:"name"`
	Email    string              `json:"email"  form:"email"`
	Phone    string              `json:"phone"  form:"phone"`
	PostCode string              `json:"post_code"  form:"post_code"`
	Address  string              `json:"address"  form:"address"`
	Status   string              `json:"status"  form:"status"`
	TotalQty int                 `json:"total_qty"  form:"total_qty"`
	SubTotal int                 `json:"sub_total"  form:"sub_total"`
	Cart     []models.Cart       `json:"carts"`
}

type UpdateTransactionRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email"  form:"email"`
	Phone    string `json:"phone"  form:"phone"`
	PostCode string `json:"post_code"  form:"post_code"`
	Address  string `json:"address"  form:"address"`
	Status   string `json:"status"  form:"status"`
	TotalQty int    `json:"total_qty" gorm:"type: int" form:"total_qty"`
	SubTotal int    `json:"sub_total" gorm:"type: int" form:"sub_total"`
}
