package repositories

import (
	"waysbean/models"

	"gorm.io/gorm"
)

type ProfiletRepository interface {
	FindProfile() ([]models.Profile, error)
	CreateProfile(profile models.Profile) (models.Profile, error)
	GetProfile(id int) (models.Profile, error)
	GetProfileByUserID(UserID int) (models.Profile, error)
}

func NewProfileRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// method array
func (r *repository) FindProfile() ([]models.Profile, error) {
	var profiles []models.Profile
	err := r.db.Preload("User").Find(&profiles).Error

	return profiles, err
}

func (r *repository) GetProfile(ID int) (models.Profile, error) {
	var profile models.Profile
	err := r.db.Preload("User").First(&profile, ID).Error

	return profile, err
}

func (r *repository) CreateProfile(profile models.Profile) (models.Profile, error) {
	err := r.db.Create(&profile).Error

	return profile, err
}

func (r *repository) GetProfileByUserID(UserID int) (models.Profile, error) {
	var profile models.Profile
	err := r.db.Where("user_id = ?", UserID).Preload("User").First(&profile).Error
	return profile, err
}
