package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/869413421/micro-chat/api/user/service/v1"
)

var _ PermissionUsecase = (*permissionUsecase)(nil)

// Permission 角色
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

// ToProtoPermission 转换为proto结构体
func (permission *Permission) ToProtoPermission() *v1.PermissionInfoResponse {
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

// PermissionRepo 注意这一行新增的 mock 数据的命令
//go:generate mockgen -source=permission.go -destination=../mocks/mrepo/permission.go -package=mrepo  PermissionRepo
type PermissionRepo interface {
	Create(context.Context, *Permission) (*Permission, error)
	Update(context.Context, *Permission) (*Permission, error)
	Delete(context.Context, uint64) (*Permission, error)
	Get(context.Context, map[string]interface{}) (*Permission, error)
	List(context.Context, map[string]interface{}, map[string]bool, int64, int64) ([]*Permission, int64, error)
}

// PermissionUsecase UserUsecase 角色业务逻辑接口
//go:generate mockgen -source=permission.go -destination=../mocks/mbiz/permission.go -package=mbiz  PermissionUsecase
type PermissionUsecase interface {
	Create(context.Context, *v1.CreatePermissionRequest) (*Permission, error)
	Update(context.Context, *v1.UpdatePermissionRequest) (*Permission, error)
	Delete(context.Context, uint64) (*Permission, error)
	Get(context.Context, *v1.GetPermissionRequest) (*Permission, error)
	List(context.Context, *v1.ListPermissionRequest) ([]*Permission, int64, error)
}

// PermissionUsecase 角色业务逻辑
type permissionUsecase struct {
	repo PermissionRepo
	log  *log.Helper
}

// NewPermissionUsecase 创建用户业务逻辑
func NewPermissionUsecase(repo PermissionRepo, logger log.Logger) PermissionUsecase {
	return &permissionUsecase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/permission"))}
}

// Create 创建角色
func (r *permissionUsecase) Create(ctx context.Context, req *v1.CreatePermissionRequest) (*Permission, error) {
	return r.repo.Create(ctx, &Permission{
		Name:     req.Name,
		Memo:     req.Memo,
		ParentID: req.ParentId,
		Path:     req.Path,
		Method:   req.Method,
	})
}

// Update 更新角色
func (r *permissionUsecase) Update(ctx context.Context, req *v1.UpdatePermissionRequest) (*Permission, error) {
	return r.repo.Update(ctx, &Permission{
		ID:       req.Id,
		Name:     req.Name,
		Memo:     req.Memo,
		ParentID: req.ParentId,
		Path:     req.Path,
		Method:   req.Method,
	})
}

// Delete 删除角色
func (r *permissionUsecase) Delete(ctx context.Context, id uint64) (*Permission, error) {
	return r.repo.Delete(ctx, id)
}

// Get 获取角色
func (r *permissionUsecase) Get(ctx context.Context, req *v1.GetPermissionRequest) (*Permission, error) {
	where := make(map[string]interface{})
	where["id = "] = req.Id
	return r.repo.Get(ctx, where)
}

// List 列表角色
func (r *permissionUsecase) List(ctx context.Context, req *v1.ListPermissionRequest) ([]*Permission, int64, error) {
	where := map[string]interface{}{}
	order := map[string]bool{
		"id": true,
	}
	return r.repo.List(ctx, where, order, req.Page, req.PageSize)
}
