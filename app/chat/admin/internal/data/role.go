package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/869413421/micro-chat/api/user/service/v1"
	"github.com/869413421/micro-chat/app/chat/admin/internal/biz"
)

var _ biz.RoleRepo = (*roleRepo)(nil)

type roleRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoleRepo(data *Data, logger log.Logger) biz.RoleRepo {
	return &roleRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/role")),
	}
}

// Get 获取角色信息
func (r *roleRepo) Get(ctx context.Context, where map[string]interface{}) (*biz.Role, error) {
	role, err := r.data.uc.GetRole(ctx, &v1.GetRoleRequest{Id: where["id"].(uint64)})
	if err != nil {
		return nil, err
	}
	return RoleResponseToBizRole(role), nil
}

// Create 创建角色
func (r *roleRepo) Create(ctx context.Context, role *biz.Role) (*biz.Role, error) {
	response, err := r.data.uc.CreateRole(ctx, &v1.CreateRoleRequest{
		Name: role.Name,
		Memo: role.Memo,
	})
	if err != nil {
		return nil, err
	}
	return RoleResponseToBizRole(response), nil
}

// Update 更新角色
func (r *roleRepo) Update(ctx context.Context, role *biz.Role) (*biz.Role, error) {
	response, err := r.data.uc.UpdateRole(ctx, &v1.UpdateRoleRequest{
		Id:   role.ID,
		Name: role.Name,
		Memo: role.Memo,
	})
	if err != nil {
		return nil, err
	}
	return RoleResponseToBizRole(response), nil
}

// Delete 删除角色
func (r *roleRepo) Delete(ctx context.Context, id uint64) (*biz.Role, error) {
	response, err := r.data.uc.DeleteRole(ctx, &v1.DeleteRoleRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return RoleResponseToBizRole(response), nil
}

// List 角色列表
func (r *roleRepo) List(ctx context.Context, where map[string]interface{}, page, pageSize int64) ([]*biz.Role, int64, error) {
	response, err := r.data.uc.ListRole(ctx, &v1.ListRoleRequest{
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		return nil, 0, err
	}

	var Roles []*biz.Role
	for _, Role := range response.Roles {
		Roles = append(Roles, RoleResponseToBizRole(Role))
	}

	return Roles, response.Total, nil
}

// RoleResponseToBizRole 类型转换
func RoleResponseToBizRole(r *v1.RoleInfoResponse) *biz.Role {
	if r == nil {
		return nil
	}
	return &biz.Role{
		ID:       r.Id,
		Name:     r.Name,
		Memo:     r.Memo,
		CreateAt: r.CreatedAt,
		UpdateAt: r.UpdatedAt,
	}
}
