package books

import (
	"api/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateBookRequestBody struct {
	Title          string       `json:"title"`
	Author         string       `json:"author"`
	Description    string       `json:"description"`
	AdditionalData models.JSONB `Gorm:"type:jsonb;serializer:json" json:"additional_data"`
}

func (h handler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	body := UpdateBookRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		_ = c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description
	book.AdditionalData = body.AdditionalData

	h.DB.Save(&book)

	c.JSON(http.StatusOK, &book)
}
