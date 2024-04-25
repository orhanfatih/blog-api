package repository

import (
	"github.com/orhanfatih/blog-api/model"
	"gorm.io/gorm"
)

type PostStore interface {
	CreatePost(post *model.Post) error
	FindPost(post *model.Post, postID int) (*model.Post, error)
	UpdatePost(post, updated *model.Post) error
	DeletePost(postId int) error
	FindPosts(limit, offset int) ([]*model.Post, error)
}

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (repo PostRepository) CreatePost(post *model.Post) error {
	tx := repo.db.Create(post)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo PostRepository) FindPost(post *model.Post, postID int) (*model.Post, error) {
	tx := repo.db.First(&post, "id = ?", postID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return post, nil
}

func (repo PostRepository) UpdatePost(post, updated *model.Post) error {
	tx := repo.db.Model(post).Updates(updated)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo PostRepository) DeletePost(postId int) error {
	tx := repo.db.Delete(&model.Post{}, "id = ?", postId)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo PostRepository) FindPosts(limit, offset int) ([]*model.Post, error) {
	var posts []*model.Post
	tx := repo.db.Limit(limit).Offset(offset).Find(&posts)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return posts, nil
}
