package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID    uuid.UUID `gorm:"column:id;type:varchar(32);primaryKey"`
	Title string    `gorm:"column:title;type:varchar(50);index:idx_title"`
}
