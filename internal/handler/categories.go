package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCategories
// @Summary Получение категорий
// @Tags categories
// @Produce json
// @Router /categories/get [get]
func (h *handler) GetCategories(c *gin.Context) {
	categories, err := h.service.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":  nil,
		"result": categories,
	})
}
