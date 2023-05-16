package data

import (
	"context"
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/869413421/micro-chat/app/user/service/internal/biz"
	"github.com/869413421/micro-chat/app/user/service/internal/data/ent/schema"
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
		log:      log.NewHelper(log.With(logger, "module", "repo/role")),
	}
}

// Create 创建角色
func (r *roleRepo) Create(ctx context.Context, bizRole *biz.Role) (*biz.Role, error) {
	var dbRole schema.Role
	result := r.data.db.Where("name = ?", bizRole.Name).First(&dbRole)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "角色已存在")
	}
	dbRole.Name = bizRole.Name
	dbRole.Memo = bizRole.Memo
	res := r.data.db.Create(&dbRole)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	return dbRole.ToBizRole(), nil
}

// Update 更新角色
func (r *roleRepo) Update(ctx context.Context, bizRole *biz.Role) (*biz.Role, error) {
	var dbRole schema.Role
	result := r.data.db.Where("name = ?", bizRole.Name).First(&dbRole)
	if result.RowsAffected > 0 {
		if dbRole.ID != bizRole.ID {
			return nil, status.Errorf(codes.AlreadyExists, "角色名重复")
		}
	} else {
		result = r.data.db.Where("id = ?", bizRole.ID).First(&dbRole)
		if result.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "更新角色不存在")
		}
	}

	dbRole.Name = bizRole.Name
	dbRole.Memo = bizRole.Memo
	res := r.data.db.Save(&dbRole)

	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	return dbRole.ToBizRole(), nil
}

// Delete 删除角色
func (r *roleRepo) Delete(ctx context.Context, id uint64) (*biz.Role, error) {
	var dbRole schema.Role
	result := r.data.db.Where("id = ?", id).First(&dbRole)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "删除角色不存在")
	}
	res := r.data.db.Delete(&dbRole)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	return dbRole.ToBizRole(), nil
}

// Get 获取角色
func (r *roleRepo) Get(ctx context.Context, where map[string]interface{}) (*biz.Role, error) {
	var dbRole schema.Role
	db := r.data.db
	for key, value := range where {
		db = db.Where(fmt.Sprintf("%s ?", key), value)
	}

	res := db.First(&dbRole)
	if res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	return dbRole.ToBizRole(), nil
}

// List 列出角色
func (r *roleRepo) List(ctx context.Context, where map[string]interface{}, order map[string]bool, page, pageSize int64) ([]*biz.Role, int64, error) {
	var dbRoles []*schema.Role
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
	res := db.Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&dbRoles)
	if res.Error != nil {
		return nil, 0, status.Errorf(codes.Internal, res.Error.Error())
	}
	var count int64
	db.Model(&schema.Role{}).Count(&count)
	var bizRoles []*biz.Role
	for _, role := range dbRoles {
		bizRoles = append(bizRoles, role.ToBizRole())
	}
	return bizRoles, count, nil
}

// Query 获取所有角色
func (r *roleRepo) Query(ctx context.Context, where map[string]interface{}) ([]*biz.Role, error) {
	var dbRoles []*schema.Role
	db := r.data.db
	for key, value := range where {
		db = db.Where(fmt.Sprintf("%s ?", key), value)
	}
	res := db.Find(&dbRoles)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	var bizRole []*biz.Role
	for _, role := range dbRoles {
		bizRole = append(bizRole, role.ToBizRole())
	}
	return bizRole, nil
}
