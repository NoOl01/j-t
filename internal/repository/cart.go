package repository

import (
	"errors"
	"gorm.io/gorm"
	"johny-tuna/internal/errs"
	"johny-tuna/internal/handler/dto"
	"johny-tuna/internal/models"
)

func (r *repository) GetCartInfo(id int64) ([]models.CartItem, error) {
	var items []models.CartItem

	if err := r.db.Preload("cart_items").Where("user_id = ?", id).Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (r *repository) UpdateCart(userID int64, cart dto.UpdateCart) error {
	var cartItem models.CartItem
	var cartModel models.Cart
	var product models.Product

	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Where("user_id = ? AND product_id = ?", userID, cart.ProductId).First(&cartItem).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return err
	}

	if err := tx.Where("id = ?", cart.ProductId).First(&product).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("user_id = ?", userID).First(&cartModel).Error; err != nil {
		tx.Rollback()
		return err
	}

	if cart.Count <= 0 {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			if err := tx.Where("user_id = ? AND product_id = ?", userID, cart.ProductId).Delete(&models.CartItem{}).Error; err != nil {
				tx.Rollback()
				return err
			}
			cartModel.Price -= float64(cartItem.Count) * product.Price
			if cartModel.Price < 0 {
				cartModel.Price = 0
			}
			if err := tx.Save(&cartModel).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			return errs.ProductZeroCount
		}
		return tx.Commit().Error
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		cartItem = models.CartItem{
			UserId:    userID,
			ProductId: cart.ProductId,
			Count:     cart.Count,
		}
		if err := tx.Create(&cartItem).Error; err != nil {
			tx.Rollback()
			return err
		}
		cartModel.Price += float64(cart.Count) * product.Price
	} else {
		diff := cart.Count - cartItem.Count
		cartItem.Count = cart.Count
		if err := tx.Save(&cartItem).Error; err != nil {
			tx.Rollback()
			return err
		}
		cartModel.Price += float64(diff) * product.Price
		if cartModel.Price < 0 {
			cartModel.Price = 0
		}
	}

	if err := tx.Save(&cartModel).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
