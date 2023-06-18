package models

import "time"

type Cart struct {
	ID            int             `json:"id" gorm:"primary_key:auto_increment"`
	UserID        int             `json:"user_id"`
	TransactionID int             `json:"transaction_id" gorm:"type: int" form:"transaction_id constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Transaction   Transaction     `json:"-" gorm:"foreignKey:TransactionID"`
	ProductID     int             `json:"product_id" gorm:"type: int" form:"product_id"`
	Product       ProductResponse `json:"products" `
	OrderQuantity int             `json:"order_quantity" gorm:"type: int" form:"order_quantity"`
	CreatedAt     time.Time       `json:"-"`
	UpdatedAt     time.Time       ` json:"-"`
}

type CartResponse struct {
	ProductID     int             `json:"product_id"`
	Product       ProductResponse `json:"products"`
	OrderQuantity int             `json:"order_quantity"`
	TransactionID int             `json:"transaction_id"`
}
type CartProductResponse struct {
	ProductID     int `json:"product_id"`
	OrderQuantity int `json:"order_quantity"`
}

func (CartResponse) TableName() string {
	return "carts"
}
