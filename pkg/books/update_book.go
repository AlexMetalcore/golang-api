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

func (handlerInit handler) UpdateBook(context *gin.Context) {
	id := context.Param("id")
	body := UpdateBookRequestBody{}

	if err := context.BindJSON(&body); err != nil {
		_ = context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	if result := handlerInit.DB.First(&book, id); result.Error != nil {
		_ = context.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

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

	handlerInit.DB.Save(&book)

	context.JSON(http.StatusOK, &book)
}
