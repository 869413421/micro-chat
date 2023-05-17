package biz

import (
	"context"
	v1 "github.com/869413421/micro-chat/api/user/service/v1"
	"github.com/869413421/micro-chat/app/user/service/internal/conf"
	"github.com/869413421/micro-chat/pkg/auth"
	password2 "github.com/869413421/micro-chat/pkg/password"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ UserUsecase = (*userUsecase)(nil)

// User 定义返回数据结构体
type User struct {
	ID       uint64
	Name     string
	Email    string
	Password string
	Status   int
	CreateAt string
	UpdateAt string
}

// ToProtoUser 转换为proto结构体
func (user *User) ToProtoUser() *v1.UserInfoResponse {
	return &v1.UserInfoResponse{
		Id:        user.ID,
		Name:      user.Name,
		Password:  user.Password,
		Email:     user.Email,
		CreatedAt: user.CreateAt,
		UpdatedAt: user.UpdateAt,
	}
}

// UserRole 定义返回数据结构体
type UserRole struct {
	ID     uint64
	UserID uint64
	RoleID uint64
}

// UserRepo 注意这一行新增的 mock 数据的命令
//go:generate mockgen -source=user.go -destination=../mocks/mrepo/user.go -package=mrepo  UserRepo
type UserRepo interface {
	CreateUser(context.Context, *User) (*User, error)
	UpdateUser(context.Context, *User) (*User, error)
	DeleteUser(context.Context, uint64) (*User, error)
	GetUser(context.Context, map[string]interface{}) (*User, error)
	ListUser(ctx context.Context, where map[string]interface{}, page, pageSize int64) ([]*User, int64, error)
}

// UserUsecase 用户业务逻辑接口
//go:generate mockgen -source=user.go -destination=../mocks/mbiz/user.go -package=mbiz  UserUsecase
type UserUsecase interface {
	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, id uint64) (*User, error)
	Get(ctx context.Context, where map[string]interface{}) (*User, error)
	List(ctx context.Context, where map[string]interface{}, page, pageSize int64) ([]*User, int64, error)
	Login(ctx context.Context, email, password string) (string, error)
	//SetUserRole(ctx context.Context, userId uint64, roleIds []uint64) ([]Role, error)
}

// UserUsecase 用户业务逻辑接口
type userUsecase struct {
	jwtSecret string
	repo      UserRepo
	roleRepo  RoleRepo
	log       *log.Helper
}

// NewUserUsecase 创建用户业务逻辑
func NewUserUsecase(repo UserRepo, roleRepo RoleRepo, ac *conf.Auth, logger log.Logger) UserUsecase {
	return &userUsecase{repo: repo, roleRepo: roleRepo, jwtSecret: ac.JwtSecret, log: log.NewHelper(log.With(logger, "module", "usecase/user"))}
}

// Login 用户登录
func (uc *userUsecase) Login(ctx context.Context, email, password string) (string, error) {
	// 获取用户信息
	user, err := uc.repo.GetUser(ctx, map[string]interface{}{"email": email})
	if err != nil {
		return "", err
	}

	// 验证密码
	if password2.CheckHash(password, user.Password) == false {
		return "", status.Errorf(codes.Unauthenticated, "密码错误")
	}
	token, err := auth.GenerateToken(user.ID, user.Name, uc.jwtSecret)
	if err != nil {
		return "", err
	}
	return token, nil
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

// List 获取用户列表
func (uc *userUsecase) List(ctx context.Context, where map[string]interface{}, page, pageSize int64) ([]*User, int64, error) {
	return uc.repo.ListUser(ctx, where, page, pageSize)
}
