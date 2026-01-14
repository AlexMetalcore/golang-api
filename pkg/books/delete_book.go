package books

import (
	"api/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handlerInit handler) DeleteBook(context *gin.Context) {
	id := context.Param("id")

	var book models.Book

	if result := handlerInit.DB.First(&book, id); result.Error != nil {
		_ = context.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	handlerInit.DB.Delete(&book)

	context.Status(http.StatusOK)
}
