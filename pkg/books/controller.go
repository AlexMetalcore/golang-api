package books

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"api/pkg/common/config"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB, loadConfig config.Config) {
	handlerInit := &handler{
		DB: db,
	}

	routes := router.Group("/books", gin.BasicAuth(gin.Accounts{
		loadConfig.Auth.Name: loadConfig.Auth.Password,
	}))

	routes.POST("/", handlerInit.AddBook)
	routes.GET("/:id", handlerInit.GetBook)
	routes.PUT("/:id", handlerInit.UpdateBook)
	routes.DELETE("/:id", handlerInit.DeleteBook)

	allBooks := router.Group("/books")
	allBooks.GET("/", handlerInit.GetBooks)
}
