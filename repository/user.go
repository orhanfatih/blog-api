package repository

import (
	"github.com/orhanfatih/blog-api/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserStore interface {
	FindUser(userID int) (*model.User, error)
	UpdateUser(userID int, updated *model.User) (*model.User, error)
	DeleteUser(user *model.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo UserRepository) FindUser(userID int) (*model.User, error) {
	var me *model.User
	tx := repo.db.First(&me, "id = ?", userID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return me, nil
}

func (repo UserRepository) UpdateUser(userID int, updated *model.User) (*model.User, error) {
	var user model.User
	tx := repo.db.Model(&user).Clauses(clause.Returning{}).Where("id = ?", userID).Updates(&updated)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (repo UserRepository) DeleteUser(user *model.User) error {
	tx := repo.db.Delete(&model.User{}, user)
	if tx.Error != nil {
		return tx.Error
	}

	tx = repo.db.Where("user_id = ?", user.ID).Delete(&model.Post{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
