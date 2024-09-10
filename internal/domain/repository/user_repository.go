package repository

import (
	"testGorm/internal/domain/entity"
)

type UserRepository interface {
	Save(user *entity.User) error
	FindByID(id string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}
