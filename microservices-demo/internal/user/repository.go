package user

import "gorm.io/gorm"

type Repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository { return &Repository{db} }

func (r *Repository) Create(user *User) error { return r.db.Create(user).Error }
func (r *Repository) GetByID(id uint) (*User, error) {
    var u User
    return &u, r.db.First(&u, id).Error
}