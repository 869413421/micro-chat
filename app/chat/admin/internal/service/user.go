package service

import (
	"context"
	"fmt"
	v1 "github.com/869413421/micro-chat/api/chat/admin/v1"
	"github.com/869413421/micro-chat/app/chat/admin/internal/biz"
	"github.com/869413421/micro-chat/pkg/auth"
)

// Login 用户登录
func (c *ChatAdmin) Login(ctx context.Context, request *v1.LoginRequest) (*v1.LoginResponse, error) {
	token, err := c.uc.Login(ctx, request)
	if err != nil {
		return nil, err
	}
	return &v1.LoginResponse{
		Token: token,
	}, nil
}

// CreateUser 创建用户
func (c *ChatAdmin) CreateUser(ctx context.Context, request *v1.CreateUserRequest) (*v1.UserInfoResponse, error) {
	claims, ok := auth.FromContext(ctx)
	fmt.Println(claims, ok)
	user, err := c.uc.Create(ctx, request)
	if err != nil {
		return nil, err
	}
	return bizUserToProtoUser(user), nil
}

// UpdateUser 更新用户
func (c *ChatAdmin) UpdateUser(ctx context.Context, request *v1.UpdateUserRequest) (*v1.UserInfoResponse, error) {
	user, err := c.uc.Update(ctx, request)
	if err != nil {
		return nil, err
	}
	return bizUserToProtoUser(user), nil
}

// DeleteUser 删除用户
func (c *ChatAdmin) DeleteUser(ctx context.Context, request *v1.DeleteUserRequest) (*v1.UserInfoResponse, error) {
	user, err := c.uc.Delete(ctx, request)
	if err != nil {
		return nil, err
	}
	return bizUserToProtoUser(user), nil
}

// UserInfo 获取用户信息
func (c *ChatAdmin) UserInfo(ctx context.Context, request *v1.UserInfoRequest) (*v1.UserInfoResponse, error) {
	user, err := c.uc.Get(ctx, request)
	if err != nil {
		return nil, err
	}
	return bizUserToProtoUser(user), nil
}

// UserList 获取用户列表
func (c *ChatAdmin) UserList(ctx context.Context, request *v1.UserListRequest) (*v1.UserListResponse, error) {
	users, total, err := c.uc.List(ctx, request)
	if err != nil {
		return nil, err
	}

	var newUsers []*v1.UserInfoResponse
	for _, user := range users {
		newUsers = append(newUsers, bizUserToProtoUser(user))
	}

	return &v1.UserListResponse{
		Total: total,
		Users: newUsers,
	}, nil
}

// SetUserRole 设置用户角色
func (c *ChatAdmin) SetUserRole(ctx context.Context, request *v1.SetUserRoleRequest) (*v1.SetUserRoleResponse, error) {
	err := c.uc.SetUserRole(ctx, request)
	if err != nil {
		return nil, err
	}
	return &v1.SetUserRoleResponse{Success: true}, nil
}

// DeleteUserRole 删除用户角色
func (c *ChatAdmin) DeleteUserRole(ctx context.Context, request *v1.DeleteUserRoleRequest) (*v1.DeleteUserRoleResponse, error) {
	err := c.uc.DeleteUserRole(ctx, request)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteUserRoleResponse{Success: true}, nil
}

// 	bizUserToProtoUser 将biz层的user转换为pb层的user
func bizUserToProtoUser(user *biz.User) *v1.UserInfoResponse {
	return &v1.UserInfoResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreateAt,
		UpdatedAt: user.UpdateAt,
	}
}
