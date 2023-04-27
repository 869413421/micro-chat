package data

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/869413421/micro-chat/app/user/service/internal/biz"
)

// userRepo 数据库操作层
type userRepo struct {
	data     *Data
	enforcer *casbin.Enforcer
	log      *log.Helper
}

// NewUserRepo 新建DAO操作仓库
func NewUserRepo(data *Data, enforcer *casbin.Enforcer, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data:     data,
		enforcer: enforcer,
		log:      log.NewHelper(logger),
	}
}

// UserToBizUser 类型转换
func (r *userRepo) UserToBizUser(u *User) *biz.User {
	if u == nil {
		return nil
	}
	return &biz.User{
		ID:       u.ID,
		Name:     u.Name,
		Password: u.Password,
		Email:    u.Email,
	}
}

// CreateUser 创建用户
func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	var user User
	result := r.data.db.Where(&biz.User{Email: u.Email}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}
	user.Email = u.Email
	user.Name = u.Name
	user.Password = u.Password
	res := r.data.db.Create(&user)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	return r.UserToBizUser(&user), nil
}

// UpdateUser 更新用户
func (r *userRepo) UpdateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	var dbUser User
	res := r.data.db.Where("id = ?", user.ID).First(&dbUser)
	if res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	if user.Name != "" {
		dbUser.Name = user.Name
	}
	if user.Password != "" {
		dbUser.Password = user.Password
	}

	if err := r.data.db.Save(&dbUser).Error; err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return r.UserToBizUser(&dbUser), nil
}

// DeleteUser 删除用户
func (r *userRepo) DeleteUser(ctx context.Context, id uint64) (*biz.User, error) {
	var dbUser User
	res := r.data.db.Where("id = ?", id).First(&dbUser)
	if res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if err := r.data.db.Delete(&dbUser).Error; err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return r.UserToBizUser(&dbUser), nil
}

// GetUser 根据条件获取用户
func (r *userRepo) GetUser(ctx context.Context, where map[string]interface{}) (*biz.User, error) {
	var dbUser User
	db := r.data.db
	for key, value := range where {
		db = db.Where(fmt.Sprintf("%s ?", key), value)
	}

	res := db.First(&dbUser)
	if res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	return r.UserToBizUser(&dbUser), nil
}

// ListUser 获取用户列表
func (r *userRepo) ListUser(ctx context.Context, where map[string]interface{}) ([]*biz.User, error) {
	var dbUsers []*User
	db := r.data.db
	for key, value := range where {
		db = db.Where(fmt.Sprintf("%s ?", key), value)
	}

	if err := db.Model(&User{}).Find(&dbUsers).Error; err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var users []*biz.User
	for _, u := range dbUsers {
		users = append(users, r.UserToBizUser(u))
	}
	return users, nil
}
