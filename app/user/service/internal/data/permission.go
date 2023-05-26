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

var _ biz.PermissionRepo = (*permissionRepo)(nil)

// permissionRepo 数据库操作层
type permissionRepo struct {
	data     *Data
	enforcer *casbin.SyncedEnforcer
	log      *log.Helper
}

// NewPermissionRepo 新建DAO操作仓库
func NewPermissionRepo(data *Data, logger log.Logger) (biz.PermissionRepo, error) {
	syncedEnforcer, err := enforcer.GetSyncedEnforcer()
	if err != nil {
		return nil, err
	}
	return &permissionRepo{
		data:     data,
		enforcer: syncedEnforcer,
		log:      log.NewHelper(log.With(logger, "module", "repo/permission")),
	}, nil
}

// Create 创建权限
func (r *permissionRepo) Create(ctx context.Context, bizPermission *biz.Permission) (*biz.Permission, error) {
	var dbPermission schema.Permission
	result := r.data.db.Where("path = ?", bizPermission.Path).Where("method = ?", bizPermission.Method).First(&dbPermission)
	fmt.Printf("rowsAffected: %d\n", result.RowsAffected)
	if result.RowsAffected > 0 {
		fmt.Printf("permission.ID: %d\n", dbPermission.ID)
		return nil, status.Errorf(codes.AlreadyExists, "权限已存在")
	}
	dbPermission.Name = bizPermission.Name
	dbPermission.Memo = bizPermission.Memo
	dbPermission.Path = bizPermission.Path
	dbPermission.Method = bizPermission.Method
	dbPermission.ParentID = bizPermission.ParentID

	res := r.data.db.Create(&dbPermission)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	return dbPermission.ToBizPermission(), nil
}

// Update 更新权限
func (r *permissionRepo) Update(ctx context.Context, bizPermission *biz.Permission) (*biz.Permission, error) {
	var dbPermission schema.Permission
	result := r.data.db.Where("path = ?", bizPermission.Path).Where("method = ?", bizPermission.Method).First(&dbPermission)
	if result.RowsAffected > 0 {
		if dbPermission.ID != bizPermission.ID {
			return nil, status.Errorf(codes.AlreadyExists, "权限名重复")
		}
	} else {
		result = r.data.db.Where("id = ?", bizPermission.ID).First(&dbPermission)
		if result.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "更新权限不存在")
		}
	}

	dbPermission.Name = bizPermission.Name
	dbPermission.Memo = bizPermission.Memo
	dbPermission.Path = bizPermission.Path
	dbPermission.Method = bizPermission.Method
	dbPermission.ParentID = bizPermission.ParentID

	res := r.data.db.Save(&dbPermission)

	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	return dbPermission.ToBizPermission(), nil
}

// Delete 删除权限
func (r *permissionRepo) Delete(ctx context.Context, id uint64) (*biz.Permission, error) {
	var dbPermission schema.Permission
	result := r.data.db.Where("id = ?", id).First(&dbPermission)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "删除权限不存在")
	}
	res := r.data.db.Delete(&dbPermission)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	return dbPermission.ToBizPermission(), nil
}

// Get 获取权限
func (r *permissionRepo) Get(ctx context.Context, where map[string]interface{}) (*biz.Permission, error) {
	var dbPermission schema.Permission
	db := r.data.db
	for key, value := range where {
		db = db.Where(fmt.Sprintf("%s ?", key), value)
	}

	res := db.First(&dbPermission)
	if res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "权限不存在")
	}
	return dbPermission.ToBizPermission(), nil
}

// List 列出权限
func (r *permissionRepo) List(ctx context.Context, where map[string]interface{}, order map[string]bool, page, pageSize int64) ([]*biz.Permission, int64, error) {
	var dbPermissions []*schema.Permission
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
	db.Model(&schema.Permission{}).Count(&count)
	var bizPermissions []*biz.Permission
	for _, permission := range dbPermissions {
		bizPermissions = append(bizPermissions, permission.ToBizPermission())
	}
	return bizPermissions, count, nil
}
