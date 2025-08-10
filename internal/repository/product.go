package repository

import "johny-tuna/internal/models"

func (r *repository) GetProductsByCategory(categoryId int64) ([]models.Product, error) {
	var products []models.Product

	err := r.db.Where("category_id = ?", categoryId).Find(&products).Error

	return products, err
}

func (r *repository) SearchProductsByName(productName string) ([]models.Product, error) {
	var products []models.Product

	err := r.db.Where("name LIKE ?", "%"+productName+"%").Find(&products).Error

	return products, err
}
