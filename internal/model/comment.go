package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID      uuid.UUID `gorm:"column:id;type:varchar(32);primaryKey"`
	PostId  uuid.UUID `gorm:"column:post_id;type:varchar(32);not null;index:idx_post_id"`
	UserId  uuid.UUID `gorm:"column:user_id;type:varchar(32);not null;index:idx_user_id"`
	Content string    `gorm:"column:content;type:varchar(500);not null"`
}

func (Comment) TableName() string {
	return "comment"
}
