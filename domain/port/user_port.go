package port

import "hexagonal/domain/entity"

type UserRepository interface {
	Save(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	GetAll() ([]*entity.User, error)
}

type UserUsecase interface {
	Register(name, email, password string) error
	Login(email, password string) (*entity.User, error)
	GetAllUsers() ([]*entity.User, error)
}
