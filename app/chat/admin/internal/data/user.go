package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/869413421/micro-chat/api/user/service/v1"
	"github.com/869413421/micro-chat/app/chat/admin/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
	}
}

// Get 获取用户信息
func (u *userRepo) Get(ctx context.Context, where map[string]interface{}) (*biz.User, error) {
	user, err := u.data.uc.GetUser(ctx, &v1.GetUserRequest{Id: where["id"].(uint64)})
	if err != nil {
		return nil, err
	}
	return UserResponseToBizUser(user), nil
}

// Create 创建用户
func (u *userRepo) Create(ctx context.Context, user *biz.User) (*biz.User, error) {
	response, err := u.data.uc.CreateUser(ctx, &v1.CreateUserRequest{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return nil, err
	}
	return UserResponseToBizUser(response), nil
}

// Update 更新用户
func (u *userRepo) Update(ctx context.Context, user *biz.User) (*biz.User, error) {
	response, err := u.data.uc.UpdateUser(ctx, &v1.UpdateUserRequest{
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return nil, err
	}
	return UserResponseToBizUser(response), nil
}

// Delete 删除用户
func (u *userRepo) Delete(ctx context.Context, id uint64) (*biz.User, error) {
	response, err := u.data.uc.DeleteUser(ctx, &v1.DeleteUserRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return UserResponseToBizUser(response), nil
}

// List 用户列表
func (u *userRepo) List(ctx context.Context, where map[string]interface{}, page, pageSize int64) ([]*biz.User, int64, error) {
	response, err := u.data.uc.ListUser(ctx, &v1.ListUserRequest{
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		return nil, 0, err
	}

	var users []*biz.User
	for _, user := range response.Users {
		users = append(users, UserResponseToBizUser(user))
	}

	return users, response.Total, nil
}

// CreateToken 用户登录
func (u *userRepo) CreateToken(ctx context.Context, email, password string) (string, error) {
	response, err := u.data.uc.CreateToken(ctx, &v1.CreateTokenRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", err
	}
	return response.Token, nil
}

// UserResponseToBizUser 类型转换
func UserResponseToBizUser(u *v1.UserInfoResponse) *biz.User {
	if u == nil {
		return nil
	}
	return &biz.User{
		ID:       u.Id,
		Name:     u.Name,
		Password: u.Password,
		Email:    u.Email,
		CreateAt: u.CreatedAt,
		UpdateAt: u.UpdatedAt,
	}
}
