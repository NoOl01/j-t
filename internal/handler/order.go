package handler

import (
	"github.com/gin-gonic/gin"
	"johny-tuna/internal/errs"
	"net/http"
	"strings"
)

// PlaceAnOrder
// @Summary Получение информации о корзине
// @Tags order
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Router /order/place [post]
func (h *handler) PlaceAnOrder(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": errs.MissingAuthToken.Error(),
		})
		return
	}

	if !strings.HasPrefix(token, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": errs.WrongAuthTokenFormat.Error(),
		})
		return
	}

	if err := h.service.PlaceAnOrder(token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}
