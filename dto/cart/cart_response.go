package cartdto

import "time"

type CartResponse struct {
	UserID        int       `json:"user_id"`
	ID            int       `json:"id"`
	ProductID     int       `json:"product_id" form:"product_id"`
	OrderQuantity int       `json:"order_quantity" form:"order_quantity"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
