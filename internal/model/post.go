package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID       uuid.UUID `gorm:"column:id;type:varchar(36);primaryKey"`
	Title    string    `gorm:"column:title;type:varchar(50);index:idx_title"`
	Content  string    `gorm:"column:content;type:text;not null"`
	AuthorId uuid.UUID `gorm:"column:author_id;index:idx_author_id;not null"`
}

func (Post) TableName() string {
	return "posts"
}
