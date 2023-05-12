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

var _ biz.PermissionRepo = (*permissionRepo)(nil)

// permissionRepo 数据库操作层
type permissionRepo struct {
	data     *Data
	enforcer *casbin.Enforcer
	log      *log.Helper
}

// NewPermissionRepo 新建DAO操作仓库
func NewPermissionRepo(data *Data, enforcer *casbin.Enforcer, logger log.Logger) biz.PermissionRepo {
	return &permissionRepo{
		data:     data,
		enforcer: enforcer,
		log:      log.NewHelper(log.With(logger, "module", "repo/permission")),
	}
}

// PermissionToBizPermission 类型转换
func (r *permissionRepo) PermissionToBizPermission(permission *Permission) *biz.Permission {
	if permission == nil {
		return nil
	}
	return &biz.Permission{
		ID:        permission.ID,
		Name:      permission.Name,
		Memo:      permission.Memo,
		Path:      permission.Path,
		Method:    permission.Method,
		ParentID:  permission.ParentID,
		ParentIDS: permission.ParentIDS,
		CreateAt:  permission.CreatedAtDate(),
		UpdateAt:  permission.UpdatedAtDate(),
	}
}

// Create 创建权限
func (r *permissionRepo) Create(ctx context.Context, bizPermission *biz.Permission) (*biz.Permission, error) {
	var permission Permission
	result := r.data.db.Where("path = ?", bizPermission.Path).Where("method = ?", bizPermission.Method).First(&permission)
	fmt.Printf("rowsAffected: %d\n", result.RowsAffected)
	if result.RowsAffected > 0 {
		fmt.Printf("permission.ID: %d\n", permission.ID)
		return nil, status.Errorf(codes.AlreadyExists, "权限已存在")
	}
	permission.Name = bizPermission.Name
	permission.Memo = bizPermission.Memo
	permission.Path = bizPermission.Path
	permission.Method = bizPermission.Method
	permission.ParentID = bizPermission.ParentID

	res := r.data.db.Create(&permission)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	return r.PermissionToBizPermission(&permission), nil
}

// Update 更新权限
func (r *permissionRepo) Update(ctx context.Context, bizPermission *biz.Permission) (*biz.Permission, error) {
	var permission Permission
	result := r.data.db.Where("path = ?", bizPermission.Path).Where("method = ?", bizPermission.Method).First(&permission)
	if result.RowsAffected > 0 {
		if permission.ID != bizPermission.ID {
			return nil, status.Errorf(codes.AlreadyExists, "权限名重复")
		}
	} else {
		result = r.data.db.Where("id = ?", bizPermission.ID).First(&permission)
		if result.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "更新权限不存在")
		}
	}

	permission.Name = bizPermission.Name
	permission.Memo = bizPermission.Memo
	permission.Path = bizPermission.Path
	permission.Method = bizPermission.Method
	permission.ParentID = bizPermission.ParentID

	res := r.data.db.Save(&permission)

	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	return r.PermissionToBizPermission(&permission), nil
}

// Delete 删除权限
func (r *permissionRepo) Delete(ctx context.Context, id uint64) (*biz.Permission, error) {
	var permission Permission
	result := r.data.db.Where("id = ?", id).First(&permission)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "删除权限不存在")
	}
	res := r.data.db.Delete(&permission)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	return r.PermissionToBizPermission(&permission), nil
}

// Get 获取权限
func (r *permissionRepo) Get(ctx context.Context, where map[string]interface{}) (*biz.Permission, error) {
	var dbPermission Permission
	db := r.data.db
	for key, value := range where {
		db = db.Where(fmt.Sprintf("%s ?", key), value)
	}

	res := db.First(&dbPermission)
	if res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "权限不存在")
	}
	return r.PermissionToBizPermission(&dbPermission), nil
}

// List 列出权限
func (r *permissionRepo) List(ctx context.Context, where map[string]interface{}, order map[string]bool, page, pageSize int64) ([]*biz.Permission, int64, error) {
	var dbPermissions []*Permission
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
	res := db.Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&dbPermissions)
	if res.Error != nil {
		return nil, 0, status.Errorf(codes.Internal, res.Error.Error())
	}
	var count int64
	db.Model(&Permission{}).Count(&count)
	var bizPermissions []*biz.Permission
	for _, permission := range dbPermissions {
		bizPermissions = append(bizPermissions, r.PermissionToBizPermission(permission))
	}
	return bizPermissions, count, nil
}
