package books

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"api/pkg/common/config"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB, c config.Config) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/books", gin.BasicAuth(gin.Accounts{
		c.Auth.Name: c.Auth.Password,
	}))

	routes.POST("/", h.AddBook)
	routes.GET("/", h.GetBooks)
	routes.GET("/:id", h.GetBook)
	routes.PUT("/:id", h.UpdateBook)
	routes.DELETE("/:id", h.DeleteBook)
}
