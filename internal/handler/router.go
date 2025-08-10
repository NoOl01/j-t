package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "johny-tuna/docs"
)

func (h *handler) Route(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		products := api.Group("/products")
		{
			products.GET("/getByCategory", h.GetProductsByCategory)
		}
		categories := api.Group("/categories")
		{
			categories.GET("/get", h.GetCategories)
		}
		search := api.Group("/search")
		{
			search.GET("/products", h.SearchProductsByName)
		}
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
