package service

import (
	"context"
	"github.com/869413421/micro-chat/app/user/service/internal/biz"

	v1 "github.com/869413421/micro-chat/api/user/service/v1"
)

// CreateRole 创建角色
func (s *UserService) CreateRole(ctx context.Context, req *v1.CreateRoleRequest) (*v1.RoleInfoResponse, error) {
	role, err := s.ruc.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return bizRoleToProtoRole(role), nil
}

// UpdateRole 更新角色
func (s *UserService) UpdateRole(ctx context.Context, req *v1.UpdateRoleRequest) (*v1.RoleInfoResponse, error) {
	role, err := s.ruc.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return bizRoleToProtoRole(role), nil
}

// DeleteRole 删除角色
func (s *UserService) DeleteRole(ctx context.Context, req *v1.DeleteRoleRequest) (*v1.RoleInfoResponse, error) {
	role, err := s.ruc.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return bizRoleToProtoRole(role), nil
}

// GetRole 获取角色
func (s *UserService) GetRole(ctx context.Context, req *v1.GetRoleRequest) (*v1.RoleInfoResponse, error) {
	role, err := s.ruc.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return bizRoleToProtoRole(role), nil
}

// ListRole 获取角色列表
func (s *UserService) ListRole(ctx context.Context, req *v1.ListRoleRequest) (*v1.ListRoleResponse, error) {
	roles, total, err := s.ruc.List(ctx, req)
	if err != nil {
		return nil, err
	}

	var list []*v1.RoleInfoResponse
	for _, role := range roles {
		list = append(list, bizRoleToProtoRole(role))
	}

	return &v1.ListRoleResponse{
		Total: total,
		Roles: list,
	}, nil
}

// bizRoleToProtoRole biz角色转proto角色
func bizRoleToProtoRole(role *biz.Role) *v1.RoleInfoResponse {
	return &v1.RoleInfoResponse{
		Id:        role.ID,
		Name:      role.Name,
		Memo:      role.Memo,
		CreatedAt: role.CreateAt,
		UpdatedAt: role.UpdateAt,
	}
}
