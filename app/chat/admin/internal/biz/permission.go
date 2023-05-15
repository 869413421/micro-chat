package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/869413421/micro-chat/api/chat/admin/v1"
)

// Permission 定义返回数据结构体
type Permission struct {
	ID        uint64
	Name      string
	Memo      string
	Path      string
	Method    string
	ParentID  uint64
	ParentIDS string
	CreateAt  string
	UpdateAt  string
}

type PermissionRepo interface {
	Get(ctx context.Context, where map[string]interface{}) (*Permission, error)
	Create(ctx context.Context, Permission *Permission) (*Permission, error)
	Update(ctx context.Context, Permission *Permission) (*Permission, error)
	Delete(ctx context.Context, id uint64) (*Permission, error)
	List(ctx context.Context, where map[string]interface{}, page, pageSize int64) ([]*Permission, int64, error)
}

type PermissionUsecase interface {
	Create(ctx context.Context, request *v1.CreatePermissionRequest) (*Permission, error)
	Update(ctx context.Context, request *v1.UpdatePermissionRequest) (*Permission, error)
	Delete(ctx context.Context, request *v1.DeletePermissionRequest) (*Permission, error)
	Get(ctx context.Context, request *v1.PermissionInfoRequest) (*Permission, error)
	List(ctx context.Context, request *v1.PermissionListRequest) ([]*Permission, int64, error)
}

type permissionUsecase struct {
	repo PermissionRepo
	log  *log.Helper
}

var _ PermissionUsecase = (*permissionUsecase)(nil)

// NewPermissionUseCase 创建权限用例
func NewPermissionUseCase(repo PermissionRepo, logger log.Logger) PermissionUsecase {
	log := log.NewHelper(log.With(logger, "module", "usecase/permission"))
	return &permissionUsecase{
		repo: repo,
		log:  log,
	}
}

// Create 创建权限
func (r *permissionUsecase) Create(ctx context.Context, request *v1.CreatePermissionRequest) (*Permission, error) {
	permission := &Permission{
		Name:     request.Name,
		Memo:     request.Memo,
		Method:   request.Method,
		Path:     request.Path,
		ParentID: request.ParentId,
	}
	return r.repo.Create(ctx, permission)
}

// Update 更新权限
func (r *permissionUsecase) Update(ctx context.Context, request *v1.UpdatePermissionRequest) (*Permission, error) {
	permission := &Permission{
		ID:       request.Id,
		Name:     request.Name,
		Memo:     request.Memo,
		Method:   request.Method,
		Path:     request.Path,
		ParentID: request.ParentId,
	}
	return r.repo.Update(ctx, permission)
}

// Delete 删除权限
func (r *permissionUsecase) Delete(ctx context.Context, request *v1.DeletePermissionRequest) (*Permission, error) {
	return r.repo.Delete(ctx, request.Id)
}

// Get 获取权限信息
func (r *permissionUsecase) Get(ctx context.Context, request *v1.PermissionInfoRequest) (*Permission, error) {
	return r.repo.Get(ctx, map[string]interface{}{
		"id": request.Id,
	})
}

// List 获取权限列表
func (r *permissionUsecase) List(ctx context.Context, request *v1.PermissionListRequest) ([]*Permission, int64, error) {
	return r.repo.List(ctx, map[string]interface{}{}, request.Page, request.PageSize)
}
