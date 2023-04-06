package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

var _ UserUsecase = (*userUsecase)(nil)

// User 定义返回数据结构体
type User struct {
	ID       uint64
	Name     string
	Email    string
	Password string
	RoleId   uint64
}

// UserRepo 注意这一行新增的 mock 数据的命令
//go:generate mockgen -source=user.go -destination=../mocks/mrepo/user.go -package=mrepo  UserRepo
type UserRepo interface {
	CreateUser(context.Context, *User) (*User, error)
	UpdateUser(context.Context, *User) (*User, error)
	DeleteUser(context.Context, uint64) (*User, error)
	GetUser(context.Context, map[string]interface{}) (*User, error)
	ListUser(context.Context, map[string]interface{}) ([]*User, error)
}

// UserUsecase 用户业务逻辑接口
//go:generate mockgen -source=server.go -destination=../mocks/mbiz/server.go -package=mbiz  UserUsecase
type UserUsecase interface {
	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, id uint64) (*User, error)
	Get(ctx context.Context, where map[string]interface{}) (*User, error)
	ListUser(ctx context.Context, where map[string]interface{}) ([]*User, error)
}

// UserUsecase 用户业务逻辑接口
type userUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase 创建用户业务逻辑
func NewUserUsecase(repo UserRepo, logger log.Logger) UserUsecase {
	return &userUsecase{repo: repo, log: log.NewHelper(logger)}
}

// Create 创建用户
func (uc *userUsecase) Create(ctx context.Context, u *User) (*User, error) {
	return uc.repo.CreateUser(ctx, u)
}

// Update 更新用户
func (uc *userUsecase) Update(ctx context.Context, u *User) (*User, error) {
	return uc.repo.UpdateUser(ctx, u)
}

// Delete 删除用户
func (uc *userUsecase) Delete(ctx context.Context, id uint64) (*User, error) {
	return uc.repo.DeleteUser(ctx, id)
}

// Get 获取用户
func (uc *userUsecase) Get(ctx context.Context, where map[string]interface{}) (*User, error) {
	return uc.repo.GetUser(ctx, where)
}

// ListUser 获取用户列表
func (uc *userUsecase) ListUser(ctx context.Context, where map[string]interface{}) ([]*User, error) {
	return uc.repo.ListUser(ctx, where)
}
