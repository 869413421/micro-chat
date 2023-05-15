package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/869413421/micro-chat/api/chat/admin/v1"
)

// Role 定义返回数据结构体
type Role struct {
	ID       uint64
	Name     string
	Memo     string
	CreateAt string
	UpdateAt string
}

type RoleRepo interface {
	Get(ctx context.Context, where map[string]interface{}) (*Role, error)
	Create(ctx context.Context, Role *Role) (*Role, error)
	Update(ctx context.Context, Role *Role) (*Role, error)
	Delete(ctx context.Context, id uint64) (*Role, error)
	List(ctx context.Context, where map[string]interface{}, page, pageSize int64) ([]*Role, int64, error)
}

type RoleUsecase interface {
	Create(ctx context.Context, request *v1.CreateRoleRequest) (*Role, error)
	Update(ctx context.Context, request *v1.UpdateRoleRequest) (*Role, error)
	Delete(ctx context.Context, request *v1.DeleteRoleRequest) (*Role, error)
	Get(ctx context.Context, request *v1.RoleInfoRequest) (*Role, error)
	List(ctx context.Context, request *v1.RoleListRequest) ([]*Role, int64, error)
}

type roleUsecase struct {
	repo RoleRepo
	log  *log.Helper
}

var _ RoleUsecase = (*roleUsecase)(nil)

// NewRoleUseCase 创建角色用例
func NewRoleUseCase(repo RoleRepo, logger log.Logger) RoleUsecase {
	log := log.NewHelper(log.With(logger, "module", "usecase/role"))
	return &roleUsecase{
		repo: repo,
		log:  log,
	}
}

// Create 创建角色
func (r *roleUsecase) Create(ctx context.Context, request *v1.CreateRoleRequest) (*Role, error) {
	role := &Role{
		Name: request.Name,
		Memo: request.Memo,
	}
	return r.repo.Create(ctx, role)
}

// Update 更新角色
func (r *roleUsecase) Update(ctx context.Context, request *v1.UpdateRoleRequest) (*Role, error) {
	role := &Role{
		ID:   request.Id,
		Name: request.Name,
		Memo: request.Memo,
	}
	return r.repo.Update(ctx, role)
}

// Delete 删除角色
func (r *roleUsecase) Delete(ctx context.Context, request *v1.DeleteRoleRequest) (*Role, error) {
	return r.repo.Delete(ctx, request.Id)
}

// Get 获取角色信息
func (r *roleUsecase) Get(ctx context.Context, request *v1.RoleInfoRequest) (*Role, error) {
	return r.repo.Get(ctx, map[string]interface{}{
		"id": request.Id,
	})
}

// List 获取角色列表
func (r *roleUsecase) List(ctx context.Context, request *v1.RoleListRequest) ([]*Role, int64, error) {
	return r.repo.List(ctx, map[string]interface{}{}, request.Page, request.PageSize)
}
