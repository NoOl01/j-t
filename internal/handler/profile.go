package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"johny-tuna/internal/errs"
	"johny-tuna/internal/handler/dto"
	"net/http"
	"strings"
)

// GetProfileInfo
// @Summary Получение информации о профиле
// @Tags profile
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Router /profile/info [get]
func (h *handler) GetProfileInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" || !strings.HasPrefix(token, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": nil,
			"error":  errs.InvalidAuthToken.Error(),
		})
		return
	}

	id, err := h.service.GetProfileIdFromToken(token)
	if err != nil {
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

	profile, err := h.service.GetProfileInfo(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"result": nil,
				"error":  errs.UserNotFound.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": profile,
		"error":  nil,
	})
}

// EditProfileEmail
// @Summary Обновление почты
// @Tags profile
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param login_body body dto.EditEmailOrLogin true "Новое значение"
// @Router /profile/update/email [post]
func (h *handler) EditProfileEmail(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" || !strings.HasPrefix(token, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": errs.InvalidAuthToken.Error(),
		})
		return
	}

	var body dto.EditEmailOrLogin
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errs.InvalidBody.Error(),
		})
		return
	}

	id, err := h.service.GetProfileIdFromToken(token)
	if err != nil {
		var httpStatus int
		if errors.Is(err, errs.InvalidAuthToken) {
			httpStatus = http.StatusUnauthorized
		} else {
			httpStatus = http.StatusInternalServerError
		}
		c.JSON(httpStatus, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.EditProfileEmail(id, body.NewValue); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": errs.UserNotFound.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}

// EditProfileLogin
// @Summary Обновление логина
// @Tags profile
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param login_body body dto.EditEmailOrLogin true "Новое значение"
// @Router /profile/update/login [post]
func (h *handler) EditProfileLogin(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" || !strings.HasPrefix(token, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": errs.InvalidAuthToken.Error(),
		})
		return
	}

	var body dto.EditEmailOrLogin
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errs.InvalidBody.Error(),
		})
		return
	}

	id, err := h.service.GetProfileIdFromToken(token)
	if err != nil {
		var httpStatus int
		if errors.Is(err, errs.InvalidAuthToken) {
			httpStatus = http.StatusUnauthorized
		} else {
			httpStatus = http.StatusInternalServerError
		}
		c.JSON(httpStatus, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.EditProfileLogin(id, body.NewValue); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": errs.UserNotFound.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}
