package models

import "time"

type Transaction struct {
	ID        int          `json:"id"`
	UserID    int          `json:"user_id" gorm:"type: int" form:"user_id"`
	User      UserResponse `json:"users" gorm:"foreignKey:UserID"`
	Name      string       `json:"name" gorm:"type: varchar(255)" form:"name"`
	Email     string       `json:"email" gorm:"type: varchar(255)" form:"email"`
	Phone     string       `json:"phone" gorm:"type: varchar(255)" form:"phone"`
	PostCode  string       `json:"post_code" gorm:"type: varchar(255)" form:"post_code"`
	Address   string       `json:"address" gorm:"type: varchar(255)" form:"address"`
	Status    string       `json:"status" gorm:"type: varchar(255)" form:"status"`
	TotalQty  int          `json:"total_qty" gorm:"type: int" form:"total_qty"`
	SubTotal  int          `json:"sub_total" gorm:"type: int" form:"sub_total"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"updated_at"`
	Cart      []Cart       `json:"carts"`
}

type TransactionResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
	PostCode string `json:"post_code"`
	Address  string `json:"address"`
	UserID   string `json:"-"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
