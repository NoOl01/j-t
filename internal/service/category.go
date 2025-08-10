package service

import "johny-tuna/internal/models"

func (s *service) GetCategories() ([]models.Category, error) {
	return s.repo.GetCategories()
}
