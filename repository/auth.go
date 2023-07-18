package repository

import (
	"github.com/orhanfatih/blog-api/models"
	"gorm.io/gorm"
)

type AuthStore interface {
	CreateUser(user *models.User) error
	FindUser(user *models.User, email string) (*models.User, error)
}

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (repo AuthRepository) CreateUser(user *models.User) error {
	tx := repo.db.Create(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo AuthRepository) FindUser(user *models.User, email string) (*models.User, error) {
	tx := repo.db.First(&user, "email = ?", email)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}
