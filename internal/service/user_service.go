package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/nanfeng/mini-blog/internal/model"
	"github.com/nanfeng/mini-blog/internal/pkg/request"
	"github.com/nanfeng/mini-blog/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.IUserRepository
}

// NewUserRepository 创建 UserService 实例对象
func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// Create 创建用户
func (svc *UserService) Create(ctx context.Context, req *request.CreateUserRequest) (uuid.UUID, error) {
	// 1.将 request.CreateUserRequest 对象转换成 model.User对象
	user := &model.User{}
	user.Username = req.Username
	user.Nickname = req.Nickname
	// 2.生成 uuid
	id, err := uuid.NewV7()
	if err != nil {
		return uuid.Nil, errors.New("主键id生成失败")
	}
	user.ID = id

	// 3.加密密码
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, errors.New("密码加密失败")
	}
	user.Password = string(password)

	// 4.调用 repository 层
	id, err = svc.repo.Create(ctx, user)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

// Delete
func (svc *UserService) Delete(ctx context.Context, id string) error {
	return svc.repo.Delete(ctx, id)
}
