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
	VerificationRegister(token string) (string, error)
	VerifyUser(token string) error
	Appeal(body dto.AppealBody) error
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}
