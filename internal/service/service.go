package service

import (
	"johny-tuna/internal/handler/dto"
	"johny-tuna/internal/models"
	"johny-tuna/internal/repository"
)

type Service interface {
	GetProductsByCategory(categoryId int64) ([]models.Product, error)
	SearchProductsByName(productName string) ([]models.Product, error)
	GetCategories() ([]models.Category, error)
	Login(loginOrEmail, password string) (string, error)
	Register(login, email, password string) error
	ResetPasswordRequest(email string) error
	VerifyOtp(email string, token int64) error
	ResetPassword(email, password string) error
	VerificationRegister(token string) (string, error)
	VerifyUser(token string) error
	Appeal(body dto.AppealBody) error
	GetProfileIdFromToken(token string) (int64, error)
	GetProfileInfo(id int64) (interface{}, error)
	EditProfileEmail(id int64, newValue string) error
	EditProfileLogin(id int64, newValue string) error
	GetAllProducts() ([]models.Product, error)
	GetCartInfo(token string) ([]models.CartItem, error)
	UpdateCart(token string, cart dto.UpdateCart) error
	PlaceAnOrder(token string) error
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}
