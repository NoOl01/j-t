package repository

import "johny-tuna/internal/models"

func (r *repository) GetProductsByCategory(categoryId int64) ([]models.Product, error) {
	var products []models.Product

	if err := r.db.Where("category_id = ?", categoryId).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *repository) SearchProductsByName(productName string) ([]models.Product, error) {
	var products []models.Product

	if err := r.db.Where("name LIKE ?", "%"+productName+"%").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *repository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product

	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
