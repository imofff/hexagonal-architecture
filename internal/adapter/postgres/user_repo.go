package postgres

import (
	"hexagonal/domain/entity"
	"hexagonal/domain/port"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) port.UserRepository {
	return &UserRepo{db: db}
}

func (r *UserRepo) Save(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepo) GetAll() ([]*entity.User, error) {
	var users []*entity.User
	err := r.db.Find(&users).Error
	return users, err
}
