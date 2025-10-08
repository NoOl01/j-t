package service

import (
	"johny-tuna/internal/handler/dto"
	"johny-tuna/internal/models"
)

func (s *service) GetCartInfo(token string) ([]models.CartItem, error) {
	claims, err := DecodeToken(token)
	if err != nil {
		return nil, err
	}

	id, err := GetIdFromToken(claims)
	if err != nil {
		return nil, err
	}

	return s.repo.GetCartInfo(id)
}

func (s *service) UpdateCart(token string, cart dto.UpdateCart) error {
	claims, err := DecodeToken(token)
	if err != nil {
		return err
	}

	id, err := GetIdFromToken(claims)
	if err != nil {
		return err
	}

	return s.repo.UpdateCart(id, cart)
}
