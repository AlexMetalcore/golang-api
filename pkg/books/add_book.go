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

	if len(body.Title) > 0 {
		book.Title = body.Title
	}

	if len(body.Author) > 0 {
		book.Author = body.Author
	}

	if len(body.Description) > 0 {
		book.Description = body.Description
	}

	if len(body.AdditionalData) > 0 {
		book.AdditionalData = body.AdditionalData
	}

	if result := h.DB.Create(&book); result.Error != nil {
		_ = c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &book)
}
