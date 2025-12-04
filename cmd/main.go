package main

import (
	"api/pkg/books"
	"api/pkg/common/config"
	"api/pkg/common/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	r := gin.Default()
	h := db.Init(c.DBUrl)

	books.RegisterRoutes(r, h)

	_ = r.Run(c.Port)
}
