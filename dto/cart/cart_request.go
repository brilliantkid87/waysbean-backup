package cartdto

type CreateCartRequest struct {
	UserID        int `json:"user_id"`
	TransactionID int `json:"transaction_id" form:"transaction_id"`
	ProductID     int `json:"product_id" form:"product_id"`
	OrderQuantity int `json:"order_quantity" form:"order_quantity"`
}

type UpdateCartRequest struct {
	ProductId     int `json:"product_id" form:"product_id"`
	OrderQuantity int `json:"order_quantity" form:"order_quantity"`
}
