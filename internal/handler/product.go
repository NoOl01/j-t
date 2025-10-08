package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetProductsByCategory
// @Summary Получение продуктов по категории
// @Tags products
// @Produce json
// @Param category_id query string true "id категории"
// @Router /products/getByCategory [get]
func (h *handler) GetProductsByCategory(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	if categoryIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "missing category_id",
			"result": nil,
		})
		return
	}

	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "invalid category_id",
			"result": nil,
		})
		return
	}

	products, err := h.service.GetProductsByCategory(categoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":  nil,
		"result": products,
	})
}

// SearchProductsByName
// @Summary Получение продуктов из поиска
// @Tags search
// @Produce json
// @Param name query string true "Имя продукта (поиск)"
// @Router /search/products [get]
func (h *handler) SearchProductsByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "missing name",
			"result": nil,
		})
		return
	}

	products, err := h.service.SearchProductsByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"result": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":  nil,
		"result": products,
	})
}

// GetAllProducts
// @Summary Получение всех продуктов
// @Tags products
// @Produce json
// @Router /products/all [get]
func (h *handler) GetAllProducts(c *gin.Context) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result": nil,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": products,
		"error":  nil,
	})
}
