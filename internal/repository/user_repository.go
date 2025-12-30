package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/nanfeng/mini-blog/internal/model"
	"github.com/nanfeng/mini-blog/internal/pkg/xerr"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(ctx context.Context, user *model.User) (uuid.UUID, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, user_map *map[string]interface{}) error
	GetById(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetUserList(ctx context.Context) (*[]model.User, error)
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
func (repo *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
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

// GetById
func (repo *UserRepository) GetById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	user, err := gorm.G[model.User](repo.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerr.ErrUserNotFount
		}
		return nil, err
	}

	return &user, nil
}

// GetUserList
func (repo *UserRepository) GetUserList(ctx context.Context) (*[]model.User, error) {
	users, err := gorm.G[model.User](repo.db).Find(ctx)
	if err != nil {
		return nil, err
	}

	return &users, nil
}
