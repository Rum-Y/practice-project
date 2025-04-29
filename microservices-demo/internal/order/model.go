package order

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID     uint
	ProductID  uint
	Quantity   int
	TotalPrice float64
	Status     string // "created", "paid", "shipped", "completed"
	OrderDate  time.Time
}
