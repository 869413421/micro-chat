package service

import (
	"context"
	v1 "github.com/869413421/micro-chat/api/chat/admin/v1"
	"github.com/869413421/micro-chat/app/chat/admin/internal/biz"
)

// CreateRole 创建角色
func (c *ChatAdmin) CreateRole(ctx context.Context, request *v1.CreateRoleRequest) (*v1.RoleInfoResponse, error) {
	role, err := c.rc.Create(ctx, request)
	if err != nil {
		return nil, err
	}
	return bizRoleToProtoRole(role), nil
}

// UpdateRole 更新角色
func (c *ChatAdmin) UpdateRole(ctx context.Context, request *v1.UpdateRoleRequest) (*v1.RoleInfoResponse, error) {
	role, err := c.rc.Update(ctx, request)
	if err != nil {
		return nil, err
	}
	return bizRoleToProtoRole(role), nil
}

// DeleteRole 删除角色
func (c *ChatAdmin) DeleteRole(ctx context.Context, request *v1.DeleteRoleRequest) (*v1.RoleInfoResponse, error) {
	role, err := c.rc.Delete(ctx, request)
	if err != nil {
		return nil, err
	}
	return bizRoleToProtoRole(role), nil
}

// RoleInfo 获取角色信息
func (c *ChatAdmin) RoleInfo(ctx context.Context, request *v1.RoleInfoRequest) (*v1.RoleInfoResponse, error) {
	role, err := c.rc.Get(ctx, request)
	if err != nil {
		return nil, err
	}
	return bizRoleToProtoRole(role), nil
}

// RoleList 角色列表
func (c *ChatAdmin) RoleList(ctx context.Context, request *v1.RoleListRequest) (*v1.RoleListResponse, error) {
	roles, total, err := c.rc.List(ctx, request)
	if err != nil {
		return nil, err
	}

	var newRoles []*v1.RoleInfoResponse
	for _, role := range roles {
		newRoles = append(newRoles, bizRoleToProtoRole(role))
	}

	return &v1.RoleListResponse{
		Total: total,
		Roles: newRoles,
	}, nil
}

// bizRoleToProtoRole biz角色转换为proto角色
func bizRoleToProtoRole(role *biz.Role) *v1.RoleInfoResponse {
	return &v1.RoleInfoResponse{
		Id:        role.ID,
		Name:      role.Name,
		Memo:      role.Memo,
		CreatedAt: role.CreateAt,
		UpdatedAt: role.UpdateAt,
	}
}
