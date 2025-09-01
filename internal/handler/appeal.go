package handler

import (
	"github.com/gin-gonic/gin"
	"johny-tuna/internal/errs"
	"johny-tuna/internal/handler/dto"
	"net/http"
)

// Appeal
// @Summary Обращение
// @Tags appeal
// @Accept json
// @Produce json
// @Param appeal_body body dto.AppealBody true "Данные обращения"
// @Router /appeal [post]
func (h *handler) Appeal(c *gin.Context) {
	var body dto.AppealBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errs.InvalidBody.Error(),
		})
		return
	}

	if body.Name == "" || body.Email == "" || body.Message == "" || body.Theme >= 5 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errs.InvalidBody.Error(),
		})
		return
	}

	if err := h.service.Appeal(body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}
