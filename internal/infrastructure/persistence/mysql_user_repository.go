package persistence

import (
	"gorm.io/gorm"
	"testGorm/internal/domain/entity"
	"testGorm/internal/domain/repository"
)

type MySQLUserRepository struct {
	db *gorm.DB
}

func NewMySQLUserRepository(db *gorm.DB) repository.UserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) Save(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *MySQLUserRepository) FindByID(id string) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, "id = ?", id).Error
	return &user, err
}
func (r *MySQLUserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, "email = ?", email).Error
	return &user, err
}
