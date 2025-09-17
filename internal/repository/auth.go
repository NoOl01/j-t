package repository

import (
	"errors"
	"gorm.io/gorm"
	"johny-tuna/internal/errs"
	"johny-tuna/internal/models"
)

const (
	LoginByUsername = iota
	LoginByEmail
)

func (r *repository) Login(loginOrEmail string, loginType int) (*models.User, error) {
	var user models.User
	switch loginType {
	case LoginByUsername:
		if err := findFirstReq("login = ?", loginOrEmail, &user, r.db); err != nil {
			return nil, err
		}
	case LoginByEmail:
		if err := findFirstReq("email = ?", loginOrEmail, &user, r.db); err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (r *repository) Register(login, email, password string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx := r.db.Begin()
			if tx.Error != nil {
				return nil, tx.Error
			}

			user = models.User{
				Login:    login,
				Email:    email,
				Password: password,
			}
			if err := tx.Create(&user).Error; err != nil {
				tx.Rollback()
				return nil, err
			}

			var points = models.Points{
				Value:  10,
				UserId: user.Id,
			}

			if err := tx.Create(&points).Error; err != nil {
				tx.Rollback()
				return nil, err
			}

			if err := tx.Commit().Error; err != nil {
				return nil, err
			}

			return &user, nil
		}
		return nil, err
	}
	return nil, errs.UserAlreadyExist
}

func (r *repository) CheckUser(email string) error {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return errs.UserAlreadyExist
}

func findFirstReq(where, loginOrEmail string, user *models.User, db *gorm.DB) error {
	if err := db.Where(where, loginOrEmail).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.UserNotFound
		}
		return err
	}
	return nil
}
