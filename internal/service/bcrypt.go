package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"johny-tuna/internal/errs"
	"johny-tuna/internal/models"
)

func EncryptPass(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CheckPass(password string, user *models.User) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return errs.WrongPassword
		}
		return err
	}
	return nil
}
