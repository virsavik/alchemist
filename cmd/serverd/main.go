package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"io.github.virsavik/alchemist/internal/controller/product"
	v1 "io.github.virsavik/alchemist/internal/handler/rest/v1"
	"io.github.virsavik/alchemist/internal/repository"
	"io.github.virsavik/alchemist/pkg/app"
	"io.github.virsavik/alchemist/pkg/db"
)

func main() {
	appCfg := app.ReadConfigFromEnv()
	if err := appCfg.IsValid(); err != nil {
		log.Fatalf("App error: %v", err)
	}

	dbConn, err := db.Connect(appCfg)
	if err != nil {
		log.Fatalf("Database connection refused: %s", err)
	}

	ctrl := product.New(repository.New(dbConn))
	handler := v1.New(ctrl)

	r := gin.Default()

	r.GET("/products", v1.HandleEndpoint(handler.GetProducts))

	r.Run(":" + appCfg.Port)
}
