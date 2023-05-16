package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/869413421/micro-chat/api/user/service/v1"
	"github.com/869413421/micro-chat/app/user/service/internal/biz"
)

var _ v1.UserServer = (*UserService)(nil)

// UserService 用户服务
type UserService struct {
	v1.UnimplementedUserServer                       // 通过嵌入UnimplementedUserServer来实现UserServer接口
	uc                         biz.UserUsecase       // 业务逻辑
	ruc                        biz.RoleUsecase       // 角色业务逻辑
	puc                        biz.PermissionUsecase // 权限业务逻辑
	log                        *log.Helper           // 日志
}

// NewUserService 创建用户服务
func NewUserService(uc biz.UserUsecase, ruc biz.RoleUsecase, puc biz.PermissionUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, ruc: ruc, puc: puc, log: log.NewHelper(log.With(logger, "module", "service/user"))}
}

// CreateUser 创建用户
func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.UserInfoResponse, error) {
	user, err := s.uc.Create(ctx, &biz.User{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}

	return bizUserToProtoUser(user), nil
}

// UpdateUser 更新用户
func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UserInfoResponse, error) {
	user, err := s.uc.Update(ctx, &biz.User{
		ID:       req.Id,
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}

	return bizUserToProtoUser(user), nil
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.UserInfoResponse, error) {
	user, err := s.uc.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return bizUserToProtoUser(user), nil
}

// GetUser 获取用户
func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.UserInfoResponse, error) {
	user, err := s.uc.Get(ctx, map[string]interface{}{"id = ": req.Id})
	if err != nil {
		return nil, err
	}

	return bizUserToProtoUser(user), nil
}

// ListUser 获取用户列表
func (s *UserService) ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserResponse, error) {
	users, total, err := s.uc.List(ctx, map[string]interface{}{}, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	newUsers := make([]*v1.UserInfoResponse, 0)
	for _, user := range users {
		newUsers = append(newUsers, bizUserToProtoUser(user))
	}

	var usersResponse = &v1.ListUserResponse{
		Users: newUsers,
		Total: total,
	}

	return usersResponse, nil
}

// SetUserRole 设置用户角色
func (s *UserService) SetUserRole(ctx context.Context, req *v1.SetUserRoleRequest) (*v1.UserInfoResponse, error) {
	user, err := s.uc.SetUserRole(ctx, req.UserId, req.RoleIds)
	if err != nil {
		return nil, err
	}

	return bizUserToProtoUser(user), nil
}

// 	bizUserToProtoUser 将biz层的user转换为pb层的user
func bizUserToProtoUser(user *biz.User) *v1.UserInfoResponse {
	return &v1.UserInfoResponse{
		Id:        user.ID,
		Name:      user.Name,
		Password:  user.Password,
		Email:     user.Email,
		CreatedAt: user.CreateAt,
		UpdatedAt: user.UpdateAt,
	}
}
