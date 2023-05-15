package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/869413421/micro-chat/api/user/service/v1"
	"github.com/869413421/micro-chat/app/chat/admin/internal/biz"
)

var _ biz.PermissionRepo = (*permissionRepo)(nil)

type permissionRepo struct {
	data *Data
	log  *log.Helper
}

func NewPermissionRepo(data *Data, logger log.Logger) biz.PermissionRepo {
	return &permissionRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/permission")),
	}
}

// Get 获取角色信息
func (r *permissionRepo) Get(ctx context.Context, where map[string]interface{}) (*biz.Permission, error) {
	permission, err := r.data.uc.GetPermission(ctx, &v1.GetPermissionRequest{Id: where["id"].(uint64)})
	if err != nil {
		return nil, err
	}
	return PermissionResponseToBizPermission(permission), nil
}

// Create 创建角色
func (r *permissionRepo) Create(ctx context.Context, permission *biz.Permission) (*biz.Permission, error) {
	response, err := r.data.uc.CreatePermission(ctx, &v1.CreatePermissionRequest{
		Name:     permission.Name,
		Memo:     permission.Memo,
		Method:   permission.Method,
		Path:     permission.Path,
		ParentId: permission.ParentID,
	})
	if err != nil {
		return nil, err
	}
	return PermissionResponseToBizPermission(response), nil
}

// Update 更新角色
func (r *permissionRepo) Update(ctx context.Context, permission *biz.Permission) (*biz.Permission, error) {
	response, err := r.data.uc.UpdatePermission(ctx, &v1.UpdatePermissionRequest{
		Id:       permission.ID,
		Name:     permission.Name,
		Memo:     permission.Memo,
		Method:   permission.Method,
		Path:     permission.Path,
		ParentId: permission.ParentID,
	})
	if err != nil {
		return nil, err
	}
	return PermissionResponseToBizPermission(response), nil
}

// Delete 删除角色
func (r *permissionRepo) Delete(ctx context.Context, id uint64) (*biz.Permission, error) {
	response, err := r.data.uc.DeletePermission(ctx, &v1.DeletePermissionRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return PermissionResponseToBizPermission(response), nil
}

// List 角色列表
func (r *permissionRepo) List(ctx context.Context, where map[string]interface{}, page, pageSize int64) ([]*biz.Permission, int64, error) {
	response, err := r.data.uc.ListPermission(ctx, &v1.ListPermissionRequest{
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		return nil, 0, err
	}

	var permissions []*biz.Permission
	for _, permission := range response.Permissions {
		permissions = append(permissions, PermissionResponseToBizPermission(permission))
	}

	return permissions, response.Total, nil
}

// PermissionResponseToBizPermission 类型转换
func PermissionResponseToBizPermission(p *v1.PermissionInfoResponse) *biz.Permission {
	if p == nil {
		return nil
	}
	return &biz.Permission{
		ID:        p.Id,
		Name:      p.Name,
		Memo:      p.Memo,
		Path:      p.Path,
		Method:    p.Method,
		ParentID:  p.ParentId,
		ParentIDS: p.ParentIds,
		CreateAt:  p.CreatedAt,
		UpdateAt:  p.UpdatedAt,
	}
}
