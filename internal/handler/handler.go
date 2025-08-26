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
}

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) Handler {
	return &handler{service: service}
}
