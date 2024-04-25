package model

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id,omitempty"`
	UserID    uint      `gorm:"not null" json:"userid,omitempty"`
	Title     string    `gorm:"uniqueIndex;not null" json:"title,omitempty"`
	Content   string    `gorm:"not null" json:"content,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdatePostRequest struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}
