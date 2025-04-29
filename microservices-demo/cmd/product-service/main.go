package main

import (
	"github.com/Rum-Y/practice-project/microservices-demo/internal/product"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("product.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&product.Product{})

	r := gin.Default()
	productHandler := product.NewHandler(product.NewRepository(db))
	productHandler.RegisterRoutes(r)

	r.Run(":8081")
}
