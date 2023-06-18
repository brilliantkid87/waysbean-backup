package profiledto

import "time"

type CreateProfileRequest struct {
	Phone        string    `json:"phone" form:"phone"`
	Address      string    `json:"address" form:"address"`
	UserID       int       `json:"user_id" form:"user_id"`
	ImageProfile string    `json:"image_profile" form:"image_profile"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

type UpdateProfileRequest struct {
	Name         string `json:"name" form:"name"`
	Phone        string `json:"phone" form:"phone" `
	Address      string `json:"address" form:"address"`
	UserID       int    `json:"user_id" form:"user_id"`
	PostCode     string `json:"post_code" form:"post_code"`
	ImageProfile string `json:"image_profile" form:"image_profile"`
}
