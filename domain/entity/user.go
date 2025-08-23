package entity

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint
	Name         string
	Email        string
	PasswordHash string
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{Name: name, Email: email, PasswordHash: string(hash)}, nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}
