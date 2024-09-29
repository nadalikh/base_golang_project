package persistence

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	//tx := r.db.Begin() // Start a transaction
	//defer func() {
	//	if r := recover(); r != nil {
	//		tx.Rollback()
	//	}
	//}()
	//
	//// Lock the entire table to prevent reads or writes
	//if err := tx.Exec("LOCK TABLES users WRITE").Error; err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//
	//// Insert the new user
	//if err := tx.Create(user).Error; err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//
	//// Unlock the table and commit the transaction
	//tx.Exec("UNLOCK TABLES")
	//tx.Commit()
	//return nil
	tx := r.db.Begin() // Start transaction
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	tx.Updates(user)
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback() // Rollback if insert fails
		return err
	}

	tx.Commit() // Commit transaction if insert succeeds
	return nil
}

func (r *MySQLUserRepository) FindByID(id string) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, "id = ?", id).Error
	return &user, err
}
func (r *MySQLUserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *MySQLUserRepository) GetAll() (*[]entity.User, error) {
	var users *[]entity.User

	// Use transaction to ensure consistency when fetching all rows
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Find(&users).Error; err != nil {
			return err
		}
		return nil
	})

	return users, err
}
func (r *MySQLUserRepository) IsEmailTaken(email string) (bool, error) {
	var user entity.User
	tx := r.db.Begin() // Start a transaction
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Use row-level locking to prevent race conditions
	err := tx.Clauses(clause.Locking{Strength: "FOR UPDATE"}).First(&user, "email = ?", email).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Commit()
			return false, nil
		}
		tx.Rollback()
		return false, err
	}

	// Commit transaction and return true if email is found
	if err := tx.Commit().Error; err != nil {
		return false, err
	}

	return true, nil
}
