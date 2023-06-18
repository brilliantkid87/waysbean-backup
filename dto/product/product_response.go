package productdto

import "time"

type ProductResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" form:"name"`
	Price       int       `json:"price" form:"price"`
	Description string    `json:"description" form:"description"`
	Stock       int       `json:"stock" form:"stock"`
	Image       string    `json:"image" form:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
