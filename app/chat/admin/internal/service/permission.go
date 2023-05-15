package service

import (
	"context"
	v1 "github.com/869413421/micro-chat/api/chat/admin/v1"
	"github.com/869413421/micro-chat/app/chat/admin/internal/biz"
)

// CreatePermission 创建角色
func (c *ChatAdmin) CreatePermission(ctx context.Context, request *v1.CreatePermissionRequest) (*v1.PermissionInfoResponse, error) {
	permission, err := c.pc.Create(ctx, request)
	if err != nil {
		return nil, err
	}
	return bizPermissionToProtoPermission(permission), nil
}

// UpdatePermission 更新角色
func (c *ChatAdmin) UpdatePermission(ctx context.Context, request *v1.UpdatePermissionRequest) (*v1.PermissionInfoResponse, error) {
	permission, err := c.pc.Update(ctx, request)
	if err != nil {
		return nil, err
	}
	return bizPermissionToProtoPermission(permission), nil
}

// DeletePermission 删除角色
func (c *ChatAdmin) DeletePermission(ctx context.Context, request *v1.DeletePermissionRequest) (*v1.PermissionInfoResponse, error) {
	permission, err := c.pc.Delete(ctx, request)
	if err != nil {
		return nil, err
	}
	return bizPermissionToProtoPermission(permission), nil
}

// PermissionInfo 获取角色信息
func (c *ChatAdmin) PermissionInfo(ctx context.Context, request *v1.PermissionInfoRequest) (*v1.PermissionInfoResponse, error) {
	permission, err := c.pc.Get(ctx, request)
	if err != nil {
		return nil, err
	}
	return bizPermissionToProtoPermission(permission), nil
}

// PermissionList 角色列表
func (c *ChatAdmin) PermissionList(ctx context.Context, request *v1.PermissionListRequest) (*v1.PermissionListResponse, error) {
	permissions, total, err := c.pc.List(ctx, request)
	if err != nil {
		return nil, err
	}

	var newPermissions []*v1.PermissionInfoResponse
	for _, permission := range permissions {
		newPermissions = append(newPermissions, bizPermissionToProtoPermission(permission))
	}

	return &v1.PermissionListResponse{
		Total:       total,
		Permissions: newPermissions,
	}, nil
}

// bizPermissionToProtoPermission biz角色转换为proto角色
func bizPermissionToProtoPermission(permission *biz.Permission) *v1.PermissionInfoResponse {
	return &v1.PermissionInfoResponse{
		Id:        permission.ID,
		Name:      permission.Name,
		Memo:      permission.Memo,
		Path:      permission.Path,
		Method:    permission.Method,
		ParentId:  permission.ParentID,
		ParentIds: permission.ParentIDS,
		CreatedAt: permission.CreateAt,
		UpdatedAt: permission.UpdateAt,
	}
}
