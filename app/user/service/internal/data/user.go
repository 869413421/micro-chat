package data

import (
	"context"
	"fmt"
	"github.com/869413421/micro-chat/pkg/enforcer"

	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/869413421/micro-chat/app/user/service/internal/biz"
	"github.com/869413421/micro-chat/app/user/service/internal/data/orm/schema"
)

// userRepo 数据库操作层
type userRepo struct {
	data     *Data
	enforcer *casbin.SyncedEnforcer
	log      *log.Helper
}

// NewUserRepo 新建DAO操作仓库
func NewUserRepo(data *Data, logger log.Logger) (biz.UserRepo, error) {
	syncedEnforcer, err := enforcer.GetSyncedEnforcer()
	if err != nil {
		return nil, err
	}
	return &userRepo{
		data:     data,
		enforcer: syncedEnforcer,
		log:      log.NewHelper(log.With(logger, "module", "data/user")),
	}, nil
}

// CreateUser 创建用户
func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	var dbUser schema.User
	result := r.data.db.Where(&biz.User{Email: u.Email}).First(&dbUser)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}
	dbUser.Email = u.Email
	dbUser.Name = u.Name
	dbUser.Password = u.Password
	res := r.data.db.Create(&dbUser)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	return dbUser.ToBizUser(), nil
}

// UpdateUser 更新用户
func (r *userRepo) UpdateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	var dbUser schema.User
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
	return dbUser.ToBizUser(), nil
}

// DeleteUser 删除用户
func (r *userRepo) DeleteUser(ctx context.Context, id uint64) (*biz.User, error) {
	var dbUser schema.User
	res := r.data.db.Where("id = ?", id).First(&dbUser)
	if res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if err := r.data.db.Delete(&dbUser).Error; err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return dbUser.ToBizUser(), nil
}

// GetUser 根据条件获取用户
func (r *userRepo) GetUser(ctx context.Context, where map[string]interface{}) (*biz.User, error) {
	var dbUser schema.User
	db := r.data.db
	for key, value := range where {
		db = db.Where(fmt.Sprintf("%s ?", key), value)
	}

	res := db.First(&dbUser)
	if res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	return dbUser.ToBizUser(), nil
}

// ListUser 获取用户列表
func (r *userRepo) ListUser(ctx context.Context, where map[string]interface{}, page, pageSize int64) ([]*biz.User, int64, error) {
	var dbUser []*schema.User
	var count int64
	db := r.data.db
	for key, value := range where {
		db = db.Where(fmt.Sprintf("%s ?", key), value)
	}
	res := db.Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&dbUser)
	if res.Error != nil {
		return nil, 0, status.Errorf(codes.Internal, res.Error.Error())
	}
	db.Model(&schema.User{}).Count(&count)
	var bizUser []*biz.User
	for _, u := range dbUser {
		bizUser = append(bizUser, u.ToBizUser())
	}
	return bizUser, count, nil
}

// QueryUserRole 查询用户角色
func (r *userRepo) QueryUserRole(ctx context.Context, where map[string]interface{}) (*biz.UserRole, error) {
	var dbUserRole schema.UserRole
	db := r.data.db
	for key, value := range where {
		db = db.Where(fmt.Sprintf("%s ?", key), value)
	}
	res := db.First(&dbUserRole)
	if res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户角色不存在")
	}
	return dbUserRole.ToBizUserRole(), nil
}

// CreateUserRole 创建用户角色
func (r *userRepo) CreateUserRole(ctx context.Context, userRole *biz.UserRole) (*biz.UserRole, error) {
	var dbUserRole schema.UserRole
	res := r.data.db.Where("user_id = ?", userRole.UserID).Where("role_id = ?", userRole.RoleID).First(&dbUserRole)
	if res.RowsAffected > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "用户角色已存在")
	}
	dbUserRole.UserID = userRole.UserID
	dbUserRole.RoleID = userRole.RoleID
	res = r.data.db.Create(&dbUserRole)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	return dbUserRole.ToBizUserRole(), nil
}

// DeleteUserRole 删除用户角色
func (r *userRepo) DeleteUserRole(ctx context.Context, id uint64) error {
	var dbUserRole schema.UserRole
	res := r.data.db.Where("id = ?", id).First(&dbUserRole)
	if res.RowsAffected == 0 {
		return status.Errorf(codes.NotFound, "用户角色不存在")
	}
	if err := r.data.db.Delete(&dbUserRole).Error; err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}
	return nil
}
