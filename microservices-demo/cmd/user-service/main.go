package main

import (
	"github.com/Rum-Y/practice-project/microservices-demo/internal/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	db.AutoMigrate(&user.User{})

	r := gin.Default()
	userHandler := user.NewHandler(user.NewRepository(db))
	userHandler.RegisterRoutes(r)
	r.Run(":8080")
}
