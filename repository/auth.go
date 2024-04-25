package repository

import (
	"github.com/orhanfatih/blog-api/model"
	"gorm.io/gorm"
)

type AuthStore interface {
	CreateUser(user *model.User) error
	FindUser(user *model.User, email string) (*model.User, error)
}

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (repo AuthRepository) CreateUser(user *model.User) error {
	tx := repo.db.Create(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo AuthRepository) FindUser(user *model.User, email string) (*model.User, error) {
	tx := repo.db.First(&user, "email = ?", email)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}
