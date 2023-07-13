package repository

import (
	"github.com/orhanfatih/blog-api/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *models.SignUpRequest) (*models.SignUpRequest, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (repo authRepository) CreateUser(user *models.SignUpRequest) (*models.SignUpRequest, error) {
	return user, nil
}
