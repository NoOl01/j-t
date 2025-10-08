package handler

import (
	"github.com/gin-gonic/gin"
	"johny-tuna/internal/service"
)

type Handler interface {
	Route(r *gin.Engine)
	GetProductsByCategory(c *gin.Context)
	SearchProductsByName(c *gin.Context)
	GetCategories(c *gin.Context)
	Login(c *gin.Context)
	Register(c *gin.Context)
	ResetPasswordRequest(c *gin.Context)
	VerifyOtp(c *gin.Context)
	ResetPassword(c *gin.Context)
	VerifyRegister(c *gin.Context)
	VerifyUser(c *gin.Context)
	GetCartInfo(c *gin.Context)
	UpdateCart(c *gin.Context)
	Appeal(c *gin.Context)
	GetProfileInfo(c *gin.Context)
	EditProfileEmail(c *gin.Context)
	EditProfileLogin(c *gin.Context)
	GetAllProducts(c *gin.Context)
	PlaceAnOrder(c *gin.Context)
}

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) Handler {
	return &handler{service: service}
}
