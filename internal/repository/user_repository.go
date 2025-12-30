package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/nanfeng/mini-blog/internal/model"
	"github.com/nanfeng/mini-blog/internal/pkg/xerr"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(ctx context.Context, user *model.User) (uuid.UUID, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id uuid.UUID, user_map *map[string]interface{}) error
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

// Delete 删除用户
func (repo *UserRepository) Delete(ctx context.Context, id string) error {
	// 执行数据库操作
	rows, err := gorm.G[model.User](repo.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return xerr.ErrInternal
	}

	if rows == 0 {
		return xerr.ErrUserNotFount
	}

	return nil
}

// Update 更新用户
func (repo *UserRepository) Update(ctx context.Context, id uuid.UUID, user_map *map[string]interface{}) error {
	// 使用原生 GORM 的 Updates 方法,支持 map[string]interface{}
	result := repo.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Updates(user_map)
	if result.Error != nil {
		return xerr.ErrInternal
	}

	if result.RowsAffected == 0 {
		return xerr.ErrUserNotFount
	}

	return nil
}
