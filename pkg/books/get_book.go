package books

import (
	"api/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		_ = c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusOK, &book)
}
