package repository

import (
	"github.com/orhanfatih/blog-api/models"
	"gorm.io/gorm"
)

type UserStore interface {
	FindUser(userID int) (*models.User, error)
	UpdateUser(userID int, updated *models.User) error
	DeleteUser(user *models.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo UserRepository) FindUser(userID int) (*models.User, error) {
	var me *models.User
	tx := repo.db.First(&me, "id = ?", userID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return me, nil
}

func (repo UserRepository) UpdateUser(userID int, updated *models.User) error {
	tx := repo.db.Model(&models.User{}).Where("id = ?", userID).Updates(
		&updated)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo UserRepository) DeleteUser(user *models.User) error {
	tx := repo.db.Delete(&models.User{}, user)
	if tx.Error != nil {
		return tx.Error
	}

	tx = repo.db.Where("user_id = ?", user.ID).Delete(&models.Post{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
