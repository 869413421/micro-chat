package service

import (
	"context"

	v1 "github.com/869413421/micro-chat/api/user/service/v1"
)

// CreatePermission 创建权限
func (s *UserService) CreatePermission(ctx context.Context, req *v1.CreatePermissionRequest) (*v1.PermissionInfoResponse, error) {
	permission, err := s.puc.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return permission.ToProtoPermission(), nil
}

// UpdatePermission 更新权限
func (s *UserService) UpdatePermission(ctx context.Context, req *v1.UpdatePermissionRequest) (*v1.PermissionInfoResponse, error) {
	permission, err := s.puc.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return permission.ToProtoPermission(), nil
}

// DeletePermission 删除权限
func (s *UserService) DeletePermission(ctx context.Context, req *v1.DeletePermissionRequest) (*v1.PermissionInfoResponse, error) {
	permission, err := s.puc.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return permission.ToProtoPermission(), nil
}

// GetPermission 获取权限
func (s *UserService) GetPermission(ctx context.Context, req *v1.GetPermissionRequest) (*v1.PermissionInfoResponse, error) {
	permission, err := s.puc.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return permission.ToProtoPermission(), nil
}

// ListPermission 获取权限列表
func (s *UserService) ListPermission(ctx context.Context, req *v1.ListPermissionRequest) (*v1.ListPermissionResponse, error) {
	permissions, total, err := s.puc.List(ctx, req)
	if err != nil {
		return nil, err
	}

	var list []*v1.PermissionInfoResponse
	for _, permission := range permissions {
		list = append(list, permission.ToProtoPermission())
	}

	return &v1.ListPermissionResponse{
		Total:       total,
		Permissions: list,
	}, nil
}
