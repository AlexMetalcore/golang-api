package books

import (
	"api/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handlerInit handler) GetBooks(context *gin.Context) {
	var books []models.Book

	if result := handlerInit.DB.Order("ID DESC").Find(&books); result.Error != nil {
		_ = context.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	context.JSON(http.StatusOK, &books)
}
