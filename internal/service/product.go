package service

import "johny-tuna/internal/models"

func (s *service) GetProductsByCategory(categoryId int64) ([]models.Product, error) {
	return s.repo.GetProductsByCategory(categoryId)
}

func (s *service) SearchProductsByName(productName string) ([]models.Product, error) {
	return s.repo.SearchProductsByName(productName)
}

func (s *service) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAllProducts()
}
