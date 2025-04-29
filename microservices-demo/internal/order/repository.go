package order

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(order *Order) error {
	return r.db.Create(order).Error
}

func (r *Repository) GetByID(id string) (*Order, error) {
	var o Order
	err := r.db.First(&o, id).Error
	return &o, err
}

func (r *Repository) UpdateStatus(id string, status string) error {
	return r.db.Model(&Order{}).
		Where("id = ?", id).
		Update("status", status).
		Error
}
