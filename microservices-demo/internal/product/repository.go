package product

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(product *Product) error {
	return r.db.Create(product).Error
}

func (r *Repository) GetByID(id string) (*Product, error) {
	var p Product
	err := r.db.First(&p, id).Error
	return &p, err
}

func (r *Repository) UpdateStock(id string, quantity int) error {
	return r.db.Model(&Product{}).
		Where("id = ?", id).
		Update("stock", gorm.Expr("stock - ?", quantity)).
		Error
}
