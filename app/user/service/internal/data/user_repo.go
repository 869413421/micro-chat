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
		log:      log.NewHelper(log.With(logger, "module", "data/user")),
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
	return UserToBizUser(&user), nil
}

// UpdateUser 更新用户
func (r *userRepo) UpdateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	var dbUser User
	res := r.data.db.Where("email = ?", user.Email).First(&dbUser)
	if res.RowsAffected > 0 {
		if dbUser.ID != user.ID {
			return nil, status.Errorf(codes.AlreadyExists, "用户名重复")
		}
	} else {
		res = r.data.db.Where("id = ?", user.ID).First(&dbUser)
		if res.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "用户不存在")
		}
	}

	if user.Name != "" {
		dbUser.Name = user.Name
	}
	if user.Password != "" {
		dbUser.Password = user.Password
	}
	if user.Email != "" {
		dbUser.Email = user.Email
	}

	if err := r.data.db.Save(&dbUser).Error; err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return UserToBizUser(&dbUser), nil
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
	return UserToBizUser(&dbUser), nil
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
	return UserToBizUser(&dbUser), nil
}

// ListUser 获取用户列表
func (r *userRepo) ListUser(ctx context.Context, where map[string]interface{}, page, pageSize int64) ([]*biz.User, int64, error) {
	var dbUser []*User
	var count int64
	db := r.data.db
	for key, value := range where {
		db = db.Where(fmt.Sprintf("%s ?", key), value)
	}
	res := db.Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&dbUser)
	if res.Error != nil {
		return nil, 0, status.Errorf(codes.Internal, res.Error.Error())
	}
	db.Model(&User{}).Count(&count)
	var bizUser []*biz.User
	for _, u := range dbUser {
		bizUser = append(bizUser, UserToBizUser(u))
	}
	return bizUser, count, nil
}

// UserToBizUser 类型转换
func UserToBizUser(u *User) *biz.User {
	if u == nil {
		return nil
	}
	return &biz.User{
		ID:       u.ID,
		Name:     u.Name,
		Password: u.Password,
		Email:    u.Email,
		CreateAt: u.CreatedAtDate(),
		UpdateAt: u.UpdatedAtDate(),
	}
}
