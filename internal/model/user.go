package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uuid.UUID `gorm:"column:id;type:varchar(32);primaryKey;" json:"id"`
	Username string    `gorm:"column:username;type:varchar(30);index:idx_username"`
	Password string    `gorm:"column:password;type:vahrchar(128);not null"`
	Nickname string    `gorm:"column:nickname"`
	Posts    []Post    `gorm:"column:posts"`
}

func (user *User) TableName() string {
	return "user"
}
