package handler

import (
	"github.com/gin-gonic/gin"
	"johny-tuna/internal/errs"
	"johny-tuna/internal/handler/dto"
	"net/http"
	"strings"
)

// GetCartInfo
// @Summary Получение информации о корзине
// @Tags cart
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Router /cart/info [get]
func (h *handler) GetCartInfo(c *gin.Context) {
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

	items, err := h.service.GetCartInfo(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": items,
		"error":  nil,
	})
}

// UpdateCart
// @Summary Получение информации о корзине
// @Description Использовать этот запрос в случае обнавления товара в корзине, для добавления нового товара или удаления товара из корзины (count = 0)
// @Tags cart
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token}"
// @Param update_cart body dto.UpdateCart true "обновление корзины"
// @Router /cart/update [post]
func (h *handler) UpdateCart(c *gin.Context) {
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

	var body dto.UpdateCart
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.UpdateCart(token, body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}
