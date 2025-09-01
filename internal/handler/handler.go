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
	VerifyRegister(c *gin.Context)
	VerifyUser(c *gin.Context)
	GetCartInfo(c *gin.Context)
	GetPoints(c *gin.Context)
	UpdateCart(c *gin.Context)
	UpdatePoints(c *gin.Context)
	Appeal(c *gin.Context)
}

type handler struct {
	service service.Service
}

func (h *handler) GetCartInfo(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *handler) GetPoints(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *handler) UpdateCart(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *handler) UpdatePoints(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewHandler(service service.Service) Handler {
	return &handler{service: service}
}
