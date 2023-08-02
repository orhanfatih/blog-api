package repository

import (
	"github.com/orhanfatih/blog-api/models"
	"gorm.io/gorm"
)

type PostStore interface {
	CreatePost(post *models.Post) error
	FindPost(post *models.Post, postID int) (*models.Post, error)
	UpdatePost(post, updated *models.Post) error
	DeletePost(postId int) error
}

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (repo PostRepository) CreatePost(post *models.Post) error {
	tx := repo.db.Create(post)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo PostRepository) FindPost(post *models.Post, postID int) (*models.Post, error) {
	tx := repo.db.First(&post, "id = ?", postID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return post, nil
}

func (repo PostRepository) UpdatePost(post, updated *models.Post) error {
	tx := repo.db.Model(post).Updates(updated)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo PostRepository) DeletePost(postId int) error {
	tx := repo.db.Delete(&models.Post{}, "id = ?", postId)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
