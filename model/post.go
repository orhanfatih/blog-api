package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

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

func (p CreatePostRequest) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Title, validation.Required, validation.Length(1, 16)),
		validation.Field(&p.Content, validation.Length(1, 160)),
	)
}
