package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/nanfeng/mini-blog/internal/model"
	"github.com/nanfeng/mini-blog/internal/pkg/request"
	"github.com/nanfeng/mini-blog/internal/pkg/util"
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
func (svc *UserService) Delete(ctx context.Context, id uuid.UUID) error {
	return svc.repo.Delete(ctx, id)
}

// Update
func (svc *UserService) Update(ctx context.Context, req *request.UpdateUserRequest) error {
	user_map := ToMap(req)
	return svc.repo.Update(ctx, req.ID, user_map)
}

func ToMap(req *request.UpdateUserRequest) *map[string]interface{} {
	user_map := make(map[string]interface{})
	if req.Nickname != nil {
		user_map["nickname"] = req.Nickname
	}

	if req.Password != nil {
		password, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			panic("密码加密失败")
		}
		user_map["password"] = password
	}

	return &user_map
}

// GetById
func (svc *UserService) GetById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return svc.repo.GetById(ctx, id)
}

// GetUserList
func (svc *UserService) GetUserList(ctx context.Context) (*[]model.User, error) {
	return svc.repo.GetUserList(ctx)
}

// PageQuery
func (svc *UserService) PageQuery(ctx context.Context, page_query *request.UserPageQuery) (*model.Page, error) {
	users, err := svc.repo.PageQuery(ctx, page_query)
	if err != nil {
		return nil, err
	}

	page := model.Page{
		Total:  len(*users),
		Page:   page_query.Page,
		Size:   page_query.Size,
		Record: users,
	}

	return &page, nil
}

// GetByUsername
func (svc *UserService) GetByUsername(ctx context.Context, username *string) (*model.User, error) {
	return svc.repo.GetByUsername(ctx, username)
}

// Login
func (svc *UserService) Login(ctx context.Context, req *request.LoginRequest) (*string, error) {
	// 1.根据用户名查询用户信息
	user, err := svc.GetByUsername(ctx, &req.Username)
	if err != nil {
		return nil, err
	}

	// 2.校验密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("password error")
	}

	// 3.生成token
	token, err := util.GenerateToken(user.ID.String())
	if err != nil {
		return nil, errors.New("token生成失败")
	}

	// 4.返回token
	return &token, nil
}
