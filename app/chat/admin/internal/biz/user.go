package biz

import (
	"context"
	v1 "github.com/869413421/micro-chat/api/chat/admin/v1"
	"github.com/go-kratos/kratos/v2/log"
)

// User 定义返回数据结构体
type User struct {
	ID       uint64
	Name     string
	Email    string
	Password string
	CreateAt string
	UpdateAt string
}

type UserRepo interface {
	Get(ctx context.Context, where map[string]interface{}) (*User, error)
	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, id uint64) (*User, error)
	List(ctx context.Context, where map[string]interface{}, page, pageSize int64) ([]*User, int64, error)
	CreateToken(ctx context.Context, email, password string) (string, error)
}

type UserUsecase interface {
	Create(ctx context.Context, request *v1.CreateUserRequest) (*User, error)
	Update(ctx context.Context, request *v1.UpdateUserRequest) (*User, error)
	Delete(ctx context.Context, request *v1.DeleteUserRequest) (*User, error)
	Get(ctx context.Context, request *v1.UserInfoRequest) (*User, error)
	List(ctx context.Context, request *v1.UserListRequest) ([]*User, int64, error)
	Login(ctx context.Context, request *v1.LoginRequest) (string, error)
}

type userUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// Login 用户登录
func (u *userUsecase) Login(ctx context.Context, request *v1.LoginRequest) (string, error) {
	return u.repo.CreateToken(ctx, request.Email, request.Password)
}

var _ UserUsecase = (*userUsecase)(nil)

// NewUserUseCase 创建用户用例
func NewUserUseCase(repo UserRepo, logger log.Logger) UserUsecase {
	l := log.NewHelper(log.With(logger, "module", "usecase/user"))
	return &userUsecase{
		repo: repo,
		log:  l,
	}
}

// Create 创建用户
func (u *userUsecase) Create(ctx context.Context, request *v1.CreateUserRequest) (*User, error) {
	user := &User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	return u.repo.Create(ctx, user)
}

// Update 更新用户
func (u *userUsecase) Update(ctx context.Context, request *v1.UpdateUserRequest) (*User, error) {
	user := &User{
		ID:       request.Id,
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	return u.repo.Update(ctx, user)
}

// Delete 删除用户
func (u *userUsecase) Delete(ctx context.Context, request *v1.DeleteUserRequest) (*User, error) {
	return u.repo.Delete(ctx, request.Id)
}

// Get 获取用户信息
func (u *userUsecase) Get(ctx context.Context, request *v1.UserInfoRequest) (*User, error) {
	return u.repo.Get(ctx, map[string]interface{}{
		"id": request.Id,
	})
}

// List 获取用户列表
func (u *userUsecase) List(ctx context.Context, request *v1.UserListRequest) ([]*User, int64, error) {
	return u.repo.List(ctx, map[string]interface{}{}, request.Page, request.PageSize)
}
