package repository

import (
	"gorm.io/gorm"
	"johny-tuna/internal/handler/dto"
	"johny-tuna/internal/models"
)

type Repository interface {
	GetProductsByCategory(categoryId int64) ([]models.Product, error)
	SearchProductsByName(productName string) ([]models.Product, error)
	GetCategories() ([]models.Category, error)
	Login(loginOrEmail string, loginType int) (*models.User, error)
	Register(login, email, password string) (*models.User, error)
	ResetPassword(email, password string) error
	GetProfileInfo(id int64) (interface{}, error)
	EditProfileEmail(id int64, newValue string) error
	EditProfileLogin(id int64, newValue string) error
	CheckUser(email string) error
	GetAllProducts() ([]models.Product, error)
	GetCartInfo(id int64) ([]models.CartItem, error)
	UpdateCart(id int64, cart dto.UpdateCart) error
	PlaceAnOrder(id int64) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
