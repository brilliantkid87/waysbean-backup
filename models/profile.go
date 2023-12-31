package models

import "time"

type Profile struct {
	ID           int          `json:"id" gorm:"primary_key:auto_increment"`
	Name         string       `json:"name" gorm:"type: varchar(255)"`
	Phone        string       `json:"phone" gorm:"type: varchar(255)"`
	PostCode     string       `json:"post_code" gorm:"type: varchar(255)"`
	Address      string       `json:"address" gorm:"type: varchar(255)"`
	UserID       int          `json:"user_id" gorm:"type: int" form:"user_id"`
	ImageProfile string       `json:"image_profile" gorm:"type: varchar(255)"`
	User         UserResponse `json:"user"`
	CreatedAt    time.Time    `json:"-"`
	UpdatedAt    time.Time    `json:"-"`
}

type ProfileResponse struct {
	Name         string `json:"name" gorm:"type: varchar(255)"`
	Phone        string `json:"phone" gorm:"type: varchar(255)"`
	Address      string `json:"address" gorm:"type: varchar(255)"`
	ImageProfile string `json:"image_profile" gorm:"type: varchar(255)"`
	UserID       int    `json:"-"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
