package usecase

import (
	"errors"
	"hexagonal/domain/entity"
	"hexagonal/domain/port"
)

type UserUsecase struct {
	repo port.UserRepository
}

func NewUserUsecase(r port.UserRepository) port.UserUsecase {
	return &UserUsecase{repo: r}
}

func (uc *UserUsecase) Register(name, email, password string) error {
	_, err := uc.repo.FindByEmail(email)
	if err == nil {
		return errors.New("email already exists")
	}
	user, err := entity.NewUser(name, email, password)
	if err != nil {
		return err
	}
	return uc.repo.Save(user)
}

func (uc *UserUsecase) Login(email, password string) (*entity.User, error) {
	user, err := uc.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if !user.ComparePassword(password) {
		return nil, errors.New("invalid password")
	}
	return user, nil
}

func (uc *UserUsecase) GetAllUsers() ([]*entity.User, error) {
	return uc.repo.GetAll()
}
