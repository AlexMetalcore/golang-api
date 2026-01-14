package books

import (
	"api/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handlerInit handler) GetBook(context *gin.Context) {
	id := context.Param("id")

	var book models.Book

	if result := handlerInit.DB.First(&book, id); result.Error != nil {
		_ = context.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	context.JSON(http.StatusOK, &book)
}
