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
		auth := api.Group("/auth")
		{
			auth.POST("/login", h.Login)
			auth.POST("/register", h.Register)
			auth.GET("/verify", h.VerifyRegister)
			auth.GET("/verify/user", h.VerifyUser)
			auth.POST("/password/reset/req", h.ResetPasswordRequest)
			auth.POST("/password/reset/verify", h.VerifyOtp)
			auth.POST("/password/reset", h.ResetPassword)
		}
		profile := api.Group("/profile")
		{
			profile.GET("/info", h.GetProfileInfo)
			profile.POST("/update/email", h.EditProfileEmail)
			profile.POST("/update/login", h.EditProfileLogin)
		}

		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		api.POST("/appeal", h.Appeal)
	}
}
