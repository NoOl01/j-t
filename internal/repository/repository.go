package repository

import (
	"gorm.io/gorm"
	"johny-tuna/internal/models"
)

type Repository interface {
	GetProductsByCategory(categoryId int64) ([]models.Product, error)
	SearchProductsByName(productName string) ([]models.Product, error)
	GetCategories() ([]models.Category, error)
	Login(loginOrEmail string, loginType int) (*models.User, error)
	Register(login, email, password string) (*models.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
