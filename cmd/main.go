package main

import (
	"api/pkg/books"
	"api/pkg/common/config"
	"api/pkg/common/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	loadConfig, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	router := gin.Default()
	dbInit := db.Init(loadConfig.DBUrl)

	books.RegisterRoutes(router, dbInit, loadConfig)

	err = router.Run(loadConfig.Port)

	if err != nil {
		return
	}
}
