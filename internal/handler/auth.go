package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"johny-tuna/internal/errs"
	"johny-tuna/internal/handler/dto"
	"net/http"
	"strings"
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
// @Summary Регистрация (запрос на верификацию почты)
// @Tags auth
// @Accept json
// @Produce json
// @Param register_body body dto.RegisterBody true "Данные для регистрации"
// @Router /auth/register [post]
func (h *handler) Register(c *gin.Context) {
	var body dto.RegisterBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errs.InvalidBody.Error(),
		})
		return
	}

	if body.Login == "" || body.Email == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errs.InvalidBody.Error(),
		})
		return
	}

	if err := h.service.Register(body.Login, body.Email, body.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}

// VerifyRegister
// @Summary Регистрация (верификация почты по OTP коду)
// @Tags auth
// @Accept json
// @Produce json
// @Param token query string true "Данные для регистрации + код"
// @Router /auth/verify [get]
func (h *handler) VerifyRegister(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error":  errs.InvalidBody.Error(),
		})
		return
	}

	jwtToken, err := h.service.VerificationRegister(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": jwtToken,
		"error":  nil,
	})
}

// VerifyUser
// @Summary Верефикация юзера по токену
// @Tags auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Router /auth/verify/user [get]
func (h *handler) VerifyUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": nil,
			"error":  errs.MissingAuthToken.Error(),
		})
		return
	}
	if !strings.HasPrefix(token, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": nil,
			"error":  errs.WrongAuthTokenFormat.Error(),
		})
		return
	}

	if err := h.service.VerifyUser(token); err != nil {
		var httpStatus int
		if errors.Is(err, errs.InvalidAuthToken) {
			httpStatus = http.StatusUnauthorized
		} else {
			httpStatus = http.StatusInternalServerError
		}
		c.JSON(httpStatus, gin.H{
			"result": nil,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
		"error":  nil,
	})
}
