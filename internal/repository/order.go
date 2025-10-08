package repository

import (
	"errors"
	"gorm.io/gorm"
	"johny-tuna/internal/models"
)

func (r *repository) PlaceAnOrder(id int64) error {
	var cartItems []models.CartItem

	if err := r.db.Where("user_id = ?", id).Find(&cartItems).Error; err != nil {
		return err
	}

	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var productIds []int64
	for _, cartItem := range cartItems {
		productIds = append(productIds, cartItem.ProductId)
	}

	var products []models.Product
	if err := r.db.Where("id IN ?", productIds).Find(&products).Error; err != nil {
		return err
	}

	priceMap := make(map[int64]float64)
	for _, p := range products {
		priceMap[p.Id] = p.Price
	}

	var totalPrice float64
	for _, cartItem := range cartItems {
		totalPrice += float64(cartItem.Count) * priceMap[cartItem.ProductId]
	}

	order := models.Order{
		UserId: id,
		Status: "N/A",
		Price:  totalPrice,
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return err
	}

	var orderItems []models.OrderItem
	for _, cartItem := range cartItems {
		orderItems = append(orderItems, models.OrderItem{
			ProductId: cartItem.ProductId,
			UserId:    id,
			OrderId:   order.Id,
			Count:     cartItem.Count,
		})
	}

	if err := tx.Create(&orderItems).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("user_id = ?", id).Delete(&models.CartItem{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	var points models.Points
	if err := tx.Where("user_id = ?", id).First(&points).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			points = models.Points{
				UserId: id,
				Value:  int(totalPrice * 0.02),
			}
			if err := tx.Create(&points).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			tx.Rollback()
			return err
		}
	} else {
		points.Value += int(totalPrice * 0.02)
		if err := tx.Save(&points).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
