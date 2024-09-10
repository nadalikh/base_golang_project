package entity

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	return &User{
		BaseModel: BaseModel{
			ID: uuid.New().String(),
		},
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (u *User) VerifyPassword(password string) error {
	test := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return test
}
