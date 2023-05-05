package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/869413421/micro-chat/api/user/service/v1"
)

var _ RoleUsecase = (*roleUsecase)(nil)

// Role 角色
type Role struct {
	ID       uint64
	Name     string
	Memo     string
	CreateAt string
	UpdateAt string
}

// RoleRepo 注意这一行新增的 mock 数据的命令
//go:generate mockgen -source=role.go -destination=../mocks/mrepo/role.go -package=mrepo  RoleRepo
type RoleRepo interface {
	Create(context.Context, *Role) (*Role, error)
	Update(context.Context, *Role) (*Role, error)
	Delete(context.Context, uint64) (*Role, error)
	Get(context.Context, map[string]interface{}) (*Role, error)
	List(context.Context, map[string]interface{}, map[string]bool, int64, int64) ([]*Role, int64, error)
}

// RoleUsecase UserUsecase 角色业务逻辑接口
//go:generate mockgen -source=role.go -destination=../mocks/mbiz/role.go -package=mbiz  RoleUsecase
type RoleUsecase interface {
	Create(context.Context, *v1.CreateRoleRequest) (*Role, error)
	Update(context.Context, *v1.UpdateRoleRequest) (*Role, error)
	Delete(context.Context, uint64) (*Role, error)
	Get(context.Context, *v1.GetRoleRequest) (*Role, error)
	List(context.Context, *v1.ListRoleRequest) ([]*Role, int64, error)
}

// roleUsecase 角色业务逻辑
type roleUsecase struct {
	repo RoleRepo
	log  *log.Helper
}

// NewRoleUsecase 创建用户业务逻辑
func NewRoleUsecase(repo RoleRepo, logger log.Logger) RoleUsecase {
	return &roleUsecase{repo: repo, log: log.NewHelper(logger)}
}

// Create 创建角色
func (r *roleUsecase) Create(ctx context.Context, req *v1.CreateRoleRequest) (*Role, error) {
	return r.repo.Create(ctx, &Role{
		Name: req.Name,
		Memo: req.Memo,
	})
}

// Update 更新角色
func (r *roleUsecase) Update(ctx context.Context, req *v1.UpdateRoleRequest) (*Role, error) {
	return r.repo.Update(ctx, &Role{
		ID:   req.Id,
		Name: req.Name,
		Memo: req.Memo,
	})
}

// Delete 删除角色
func (r *roleUsecase) Delete(ctx context.Context, id uint64) (*Role, error) {
	return r.repo.Delete(ctx, id)
}

// Get 获取角色
func (r *roleUsecase) Get(ctx context.Context, req *v1.GetRoleRequest) (*Role, error) {
	where := make(map[string]interface{})
	where["id = "] = req.Id
	return r.repo.Get(ctx, where)
}

// List 列表角色
func (r *roleUsecase) List(ctx context.Context, req *v1.ListRoleRequest) ([]*Role, int64, error) {
	where := map[string]interface{}{}
	order := map[string]bool{
		"id": true,
	}
	return r.repo.List(ctx, where, order, req.Page, req.PageSize)
}
