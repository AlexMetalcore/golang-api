package books

import (
	"api/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddBookRequestBody struct {
	Title          string       `json:"title"`
	Author         string       `json:"author"`
	Description    string       `json:"description"`
	AdditionalData models.JSONB `Gorm:"type:jsonb;serializer:json" json:"additional_data"`
}

func (h handler) AddBook(c *gin.Context) {
	body := AddBookRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description
	book.AdditionalData = body.AdditionalData

	if result := h.DB.Create(&book); result.Error != nil {
		_ = c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &book)
}
