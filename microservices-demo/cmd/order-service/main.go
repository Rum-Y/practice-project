package main

import (
	"github.com/Rum-Y/practice-project/microservices-demo/internal/order"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("order.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&order.Order{})

	r := gin.Default()
	orderHandler := order.NewHandler(order.NewRepository(db))
	orderHandler.RegisterRoutes(r)

	r.Run(":8082")
}
