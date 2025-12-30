package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/nanfeng/mini-blog/internal/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(ctx context.Context, user *model.User) (uuid.UUID, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) Create(ctx context.Context, user *model.User) (uuid.UUID, error) {
	// 执行数据库操作
	if err := gorm.G[model.User](repo.db).Create(ctx, user); err != nil {
		return uuid.Nil, err
	}

	// 返回结果
	return user.ID, nil
}
