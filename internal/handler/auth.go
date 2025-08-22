package handler

import (
	"github.com/gin-gonic/gin"
	"johny-tuna/internal/errs"
	"johny-tuna/internal/handler/dto"
	"net/http"
)

// Login
// @Summary Вход в аккаунт
// @Tags auth
// @Accept json
// @Produce json
// @Param login_body body dto.LoginBody true "Данные для входа"
// @Router /auth/login [post]
func (h *handler) Login(c *gin.Context) {
	var body dto.LoginBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error":  errs.InvalidBody.Error(),
		})
		return
	}

	if body.LoginOrEmail == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error":  errs.InvalidBody.Error(),
		})
		return
	}

	token, err := h.service.Login(body.LoginOrEmail, body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": token,
		"error":  nil,
	})
}

// Register
// @Summary Регистрация
// @Tags auth
// @Accept json
// @Produce json
// @Param register_body body dto.RegisterBody true "Данные для регистрации"
// @Router /auth/register [post]
func (h *handler) Register(c *gin.Context) {
	var body dto.RegisterBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error":  errs.InvalidBody.Error(),
		})
		return
	}

	if body.Login == "" || body.Email == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error":  errs.InvalidBody.Error(),
		})
		return
	}

	token, err := h.service.Register(body.Login, body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": token,
		"error":  nil,
	})
}
