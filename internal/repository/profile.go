package repository

import (
	"gorm.io/gorm"
	"johny-tuna/internal/models"
)

func (r *repository) GetProfileInfo(id int64) (interface{}, error) {
	var user models.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) EditProfileEmail(id int64, newValue string) error {
	result := r.db.Model(&models.User{}).Where("id = ?", id).Update("email", newValue)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *repository) EditProfileLogin(id int64, newValue string) error {
	result := r.db.Model(&models.User{}).Where("id = ?", id).Update("login", newValue)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
