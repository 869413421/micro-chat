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

var _ biz.RoleRepo = (*roleRepo)(nil)

// roleRepo 数据库操作层
type roleRepo struct {
	data     *Data
	enforcer *casbin.Enforcer
	log      *log.Helper
}

// NewRoleRepo 新建DAO操作仓库
func NewRoleRepo(data *Data, enforcer *casbin.Enforcer, logger log.Logger) biz.RoleRepo {
	return &roleRepo{
		data:     data,
		enforcer: enforcer,
		log:      log.NewHelper(logger),
	}
}

// RoleToBizRole 类型转换
func (r *roleRepo) RoleToBizRole(role *Role) *biz.Role {
	if role == nil {
		return nil
	}
	return &biz.Role{
		ID:   role.ID,
		Name: role.Name,
		Memo: role.Memo,
	}
}

// Create 创建角色
func (r *roleRepo) Create(ctx context.Context, bizRole *biz.Role) (*biz.Role, error) {
	var role Role
	result := r.data.db.Where("name = ?", bizRole.Name).First(&role)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "角色已存在")
	}
	role.Name = bizRole.Name
	role.Memo = bizRole.Memo
	res := r.data.db.Create(&role)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	return r.RoleToBizRole(&role), nil
}

// Update 更新角色
func (r *roleRepo) Update(ctx context.Context, bizRole *biz.Role) (*biz.Role, error) {
	var role Role
	result := r.data.db.Where("name = ?", bizRole.Name).First(&role)
	if result.RowsAffected > 0 {
		if role.ID != bizRole.ID {
			return nil, status.Errorf(codes.AlreadyExists, "角色名重复")
		}
	}

	result = r.data.db.Where("id = ?", bizRole.ID).First(&role)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "更新角色不存在")
	}
	role.Name = bizRole.Name
	role.Memo = bizRole.Memo
	res := r.data.db.Save(&role)

	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	return r.RoleToBizRole(&role), nil
}

// Delete 删除角色
func (r *roleRepo) Delete(ctx context.Context, id uint64) (*biz.Role, error) {
	var role Role
	result := r.data.db.Where("id = ?", id).First(&role)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "删除角色不存在")
	}
	res := r.data.db.Delete(&role)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	return r.RoleToBizRole(&role), nil
}

// Get 获取角色
func (r *roleRepo) Get(ctx context.Context, where map[string]interface{}) (*biz.Role, error) {
	var dbRole Role
	db := r.data.db
	for key, value := range where {
		db = db.Where(fmt.Sprintf("%s ?", key), value)
	}

	res := db.First(&dbRole)
	if res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	return r.RoleToBizRole(&dbRole), nil
}

// List 列出角色
func (r *roleRepo) List(ctx context.Context, where map[string]interface{}, order map[string]bool, offset, limit int64) ([]*biz.Role, int64, error) {
	var dbRole []*Role
	db := r.data.db
	for key, value := range where {
		db = db.Where(fmt.Sprintf("%s ?", key), value)
	}
	for key, value := range order {
		if value {
			db = db.Order(fmt.Sprintf("%s desc", key))
		} else {
			db = db.Order(fmt.Sprintf("%s asc", key))
		}
	}
	res := db.Offset(int(offset)).Limit(int(limit)).Find(&dbRole)
	if res.Error != nil {
		return nil, 0, status.Errorf(codes.Internal, res.Error.Error())
	}
	var count int64
	db.Model(&Role{}).Count(&count)
	var bizRole []*biz.Role
	for _, role := range dbRole {
		bizRole = append(bizRole, r.RoleToBizRole(role))
	}
	return bizRole, count, nil
}
