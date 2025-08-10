package repository

import "johny-tuna/internal/models"

func (r *repository) GetCategories() ([]models.Category, error) {
	var categories []models.Category

	err := r.db.Find(&categories).Error

	return categories, err
}
