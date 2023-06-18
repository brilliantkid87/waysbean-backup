package productdto

type CreateProductRequest struct {
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Stock       int    `json:"stock" form:"stock"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
}

type UpdateProductRequest struct {
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Stock       int    `json:"stock" form:"stock"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
}
