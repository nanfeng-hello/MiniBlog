package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"column:id;type:varchar(36);primaryKey;" json:"ID"`
	Username string    `gorm:"column:username;type:varchar(30);index:idx_username"`
	Password string    `gorm:"column:password;type:varchar(128);not null" json:"-"`
	Nickname string    `gorm:"column:nickname"`
	Posts    []Post    `gorm:"foreignKey:AuthorId"`
}

func (user *User) TableName() string {
	return "user"
}
