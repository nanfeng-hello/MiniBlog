package request

import "github.com/google/uuid"

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=4,max=30"`
	Password string `json:"password" binding:"required,min=6,max=20"`
	Nickname string `json:"nickname" binding:"required,min=1,max=20"`
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	ID       uuid.UUID `json:"id" binding:"required,uuid"`
	Password *string   `json:"password" binding:"min=6,max=20"`
	Nickname *string   `json:"nickname" binding:"min=1,max=20"`
}

type UserPageQuery struct {
	Page     uint    `json:"page" binding:"required,page"`
	Size     uint    `json:"size" binding:"required,size"`
	Nickname *string `json:"nickname,omitempty" binding:"nickname"`
}
