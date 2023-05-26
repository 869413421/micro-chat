// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.0
// source: v1/admin.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ChatAdmin_Login_FullMethodName            = "/api.chat.admin.v1.ChatAdmin/Login"
	ChatAdmin_CreateUser_FullMethodName       = "/api.chat.admin.v1.ChatAdmin/CreateUser"
	ChatAdmin_UpdateUser_FullMethodName       = "/api.chat.admin.v1.ChatAdmin/UpdateUser"
	ChatAdmin_DeleteUser_FullMethodName       = "/api.chat.admin.v1.ChatAdmin/DeleteUser"
	ChatAdmin_UserInfo_FullMethodName         = "/api.chat.admin.v1.ChatAdmin/UserInfo"
	ChatAdmin_UserList_FullMethodName         = "/api.chat.admin.v1.ChatAdmin/UserList"
	ChatAdmin_CreateRole_FullMethodName       = "/api.chat.admin.v1.ChatAdmin/CreateRole"
	ChatAdmin_UpdateRole_FullMethodName       = "/api.chat.admin.v1.ChatAdmin/UpdateRole"
	ChatAdmin_DeleteRole_FullMethodName       = "/api.chat.admin.v1.ChatAdmin/DeleteRole"
	ChatAdmin_RoleInfo_FullMethodName         = "/api.chat.admin.v1.ChatAdmin/RoleInfo"
	ChatAdmin_RoleList_FullMethodName         = "/api.chat.admin.v1.ChatAdmin/RoleList"
	ChatAdmin_CreatePermission_FullMethodName = "/api.chat.admin.v1.ChatAdmin/CreatePermission"
	ChatAdmin_UpdatePermission_FullMethodName = "/api.chat.admin.v1.ChatAdmin/UpdatePermission"
	ChatAdmin_DeletePermission_FullMethodName = "/api.chat.admin.v1.ChatAdmin/DeletePermission"
	ChatAdmin_PermissionInfo_FullMethodName   = "/api.chat.admin.v1.ChatAdmin/PermissionInfo"
	ChatAdmin_PermissionList_FullMethodName   = "/api.chat.admin.v1.ChatAdmin/PermissionList"
	ChatAdmin_SetUserRole_FullMethodName      = "/api.chat.admin.v1.ChatAdmin/SetUserRole"
	ChatAdmin_DeleteUserRole_FullMethodName   = "/api.chat.admin.v1.ChatAdmin/DeleteUserRole"
)

// ChatAdminClient is the client API for ChatAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatAdminClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	UserList(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error)
	CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*RoleInfoResponse, error)
	UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*RoleInfoResponse, error)
	DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*RoleInfoResponse, error)
	RoleInfo(ctx context.Context, in *RoleInfoRequest, opts ...grpc.CallOption) (*RoleInfoResponse, error)
	RoleList(ctx context.Context, in *RoleListRequest, opts ...grpc.CallOption) (*RoleListResponse, error)
	CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*PermissionInfoResponse, error)
	UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*PermissionInfoResponse, error)
	DeletePermission(ctx context.Context, in *DeletePermissionRequest, opts ...grpc.CallOption) (*PermissionInfoResponse, error)
	PermissionInfo(ctx context.Context, in *PermissionInfoRequest, opts ...grpc.CallOption) (*PermissionInfoResponse, error)
	PermissionList(ctx context.Context, in *PermissionListRequest, opts ...grpc.CallOption) (*PermissionListResponse, error)
	SetUserRole(ctx context.Context, in *SetUserRoleRequest, opts ...grpc.CallOption) (*SetUserRoleResponse, error)
	DeleteUserRole(ctx context.Context, in *DeleteUserRoleRequest, opts ...grpc.CallOption) (*DeleteUserRoleResponse, error)
}

type chatAdminClient struct {
	cc grpc.ClientConnInterface
}

func NewChatAdminClient(cc grpc.ClientConnInterface) ChatAdminClient {
	return &chatAdminClient{cc}
}

func (c *chatAdminClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, ChatAdmin_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAdminClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, ChatAdmin_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAdminClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, ChatAdmin_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAdminClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, ChatAdmin_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAdminClient) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, ChatAdmin_UserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAdminClient) UserList(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, ChatAdmin_UserList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAdminClient) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*RoleInfoResponse, error) {
	out := new(RoleInfoResponse)
	err := c.cc.Invoke(ctx, ChatAdmin_CreateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAdminClient) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*RoleInfoResponse, error) {
	out := new(RoleInfoResponse)
	err := c.cc.Invoke(ctx, ChatAdmin_UpdateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAdminClient) DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*RoleInfoResponse, error) {
	out := new(RoleInfoResponse)
	err := c.cc.Invoke(ctx, ChatAdmin_DeleteRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAdminClient) RoleInfo(ctx context.Context, in *RoleInfoRequest, opts ...grpc.CallOption) (*RoleInfoResponse, error) {
	out := new(RoleInfoResponse)
	err := c.cc.Invoke(ctx, ChatAdmin_RoleInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAdminClient) RoleList(ctx context.Context, in *RoleListRequest, opts ...grpc.CallOption) (*RoleListResponse, error) {
	out := new(RoleListResponse)
	err := c.cc.Invoke(ctx, ChatAdmin_RoleList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAdminClient) CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*PermissionInfoResponse, error) {
	out := new(PermissionInfoResponse)
	err := c.cc.Invoke(ctx, ChatAdmin_CreatePermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAdminClient) UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*PermissionInfoResponse, error) {
	out := new(PermissionInfoResponse)
	err := c.cc.Invoke(ctx, ChatAdmin_UpdatePermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAdminClient) DeletePermission(ctx context.Context, in *DeletePermissionRequest, opts ...grpc.CallOption) (*PermissionInfoResponse, error) {
	out := new(PermissionInfoResponse)
	err := c.cc.Invoke(ctx, ChatAdmin_DeletePermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAdminClient) PermissionInfo(ctx context.Context, in *PermissionInfoRequest, opts ...grpc.CallOption) (*PermissionInfoResponse, error) {
	out := new(PermissionInfoResponse)
	err := c.cc.Invoke(ctx, ChatAdmin_PermissionInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAdminClient) PermissionList(ctx context.Context, in *PermissionListRequest, opts ...grpc.CallOption) (*PermissionListResponse, error) {
	out := new(PermissionListResponse)
	err := c.cc.Invoke(ctx, ChatAdmin_PermissionList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAdminClient) SetUserRole(ctx context.Context, in *SetUserRoleRequest, opts ...grpc.CallOption) (*SetUserRoleResponse, error) {
	out := new(SetUserRoleResponse)
	err := c.cc.Invoke(ctx, ChatAdmin_SetUserRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAdminClient) DeleteUserRole(ctx context.Context, in *DeleteUserRoleRequest, opts ...grpc.CallOption) (*DeleteUserRoleResponse, error) {
	out := new(DeleteUserRoleResponse)
	err := c.cc.Invoke(ctx, ChatAdmin_DeleteUserRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatAdminServer is the server API for ChatAdmin service.
// All implementations must embed UnimplementedChatAdminServer
// for forward compatibility
type ChatAdminServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*UserInfoResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UserInfoResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*UserInfoResponse, error)
	UserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error)
	UserList(context.Context, *UserListRequest) (*UserListResponse, error)
	CreateRole(context.Context, *CreateRoleRequest) (*RoleInfoResponse, error)
	UpdateRole(context.Context, *UpdateRoleRequest) (*RoleInfoResponse, error)
	DeleteRole(context.Context, *DeleteRoleRequest) (*RoleInfoResponse, error)
	RoleInfo(context.Context, *RoleInfoRequest) (*RoleInfoResponse, error)
	RoleList(context.Context, *RoleListRequest) (*RoleListResponse, error)
	CreatePermission(context.Context, *CreatePermissionRequest) (*PermissionInfoResponse, error)
	UpdatePermission(context.Context, *UpdatePermissionRequest) (*PermissionInfoResponse, error)
	DeletePermission(context.Context, *DeletePermissionRequest) (*PermissionInfoResponse, error)
	PermissionInfo(context.Context, *PermissionInfoRequest) (*PermissionInfoResponse, error)
	PermissionList(context.Context, *PermissionListRequest) (*PermissionListResponse, error)
	SetUserRole(context.Context, *SetUserRoleRequest) (*SetUserRoleResponse, error)
	DeleteUserRole(context.Context, *DeleteUserRoleRequest) (*DeleteUserRoleResponse, error)
	mustEmbedUnimplementedChatAdminServer()
}

// UnimplementedChatAdminServer must be embedded to have forward compatible implementations.
type UnimplementedChatAdminServer struct {
}

func (UnimplementedChatAdminServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedChatAdminServer) CreateUser(context.Context, *CreateUserRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedChatAdminServer) UpdateUser(context.Context, *UpdateUserRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedChatAdminServer) DeleteUser(context.Context, *DeleteUserRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedChatAdminServer) UserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedChatAdminServer) UserList(context.Context, *UserListRequest) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (UnimplementedChatAdminServer) CreateRole(context.Context, *CreateRoleRequest) (*RoleInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedChatAdminServer) UpdateRole(context.Context, *UpdateRoleRequest) (*RoleInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedChatAdminServer) DeleteRole(context.Context, *DeleteRoleRequest) (*RoleInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedChatAdminServer) RoleInfo(context.Context, *RoleInfoRequest) (*RoleInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleInfo not implemented")
}
func (UnimplementedChatAdminServer) RoleList(context.Context, *RoleListRequest) (*RoleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleList not implemented")
}
func (UnimplementedChatAdminServer) CreatePermission(context.Context, *CreatePermissionRequest) (*PermissionInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePermission not implemented")
}
func (UnimplementedChatAdminServer) UpdatePermission(context.Context, *UpdatePermissionRequest) (*PermissionInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePermission not implemented")
}
func (UnimplementedChatAdminServer) DeletePermission(context.Context, *DeletePermissionRequest) (*PermissionInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePermission not implemented")
}
func (UnimplementedChatAdminServer) PermissionInfo(context.Context, *PermissionInfoRequest) (*PermissionInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionInfo not implemented")
}
func (UnimplementedChatAdminServer) PermissionList(context.Context, *PermissionListRequest) (*PermissionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionList not implemented")
}
func (UnimplementedChatAdminServer) SetUserRole(context.Context, *SetUserRoleRequest) (*SetUserRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserRole not implemented")
}
func (UnimplementedChatAdminServer) DeleteUserRole(context.Context, *DeleteUserRoleRequest) (*DeleteUserRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserRole not implemented")
}
func (UnimplementedChatAdminServer) mustEmbedUnimplementedChatAdminServer() {}

// UnsafeChatAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatAdminServer will
// result in compilation errors.
type UnsafeChatAdminServer interface {
	mustEmbedUnimplementedChatAdminServer()
}

func RegisterChatAdminServer(s grpc.ServiceRegistrar, srv ChatAdminServer) {
	s.RegisterService(&ChatAdmin_ServiceDesc, srv)
}

func _ChatAdmin_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAdminServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatAdmin_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAdminServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAdmin_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAdminServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatAdmin_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAdminServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAdmin_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAdminServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatAdmin_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAdminServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAdmin_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAdminServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatAdmin_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAdminServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAdmin_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAdminServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatAdmin_UserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAdminServer).UserInfo(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAdmin_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAdminServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatAdmin_UserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAdminServer).UserList(ctx, req.(*UserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAdmin_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAdminServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatAdmin_CreateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAdminServer).CreateRole(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAdmin_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAdminServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatAdmin_UpdateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAdminServer).UpdateRole(ctx, req.(*UpdateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAdmin_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAdminServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatAdmin_DeleteRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAdminServer).DeleteRole(ctx, req.(*DeleteRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAdmin_RoleInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAdminServer).RoleInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatAdmin_RoleInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAdminServer).RoleInfo(ctx, req.(*RoleInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAdmin_RoleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAdminServer).RoleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatAdmin_RoleList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAdminServer).RoleList(ctx, req.(*RoleListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAdmin_CreatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAdminServer).CreatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatAdmin_CreatePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAdminServer).CreatePermission(ctx, req.(*CreatePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAdmin_UpdatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAdminServer).UpdatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatAdmin_UpdatePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAdminServer).UpdatePermission(ctx, req.(*UpdatePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAdmin_DeletePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAdminServer).DeletePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatAdmin_DeletePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAdminServer).DeletePermission(ctx, req.(*DeletePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAdmin_PermissionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAdminServer).PermissionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatAdmin_PermissionInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAdminServer).PermissionInfo(ctx, req.(*PermissionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAdmin_PermissionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAdminServer).PermissionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatAdmin_PermissionList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAdminServer).PermissionList(ctx, req.(*PermissionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAdmin_SetUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAdminServer).SetUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatAdmin_SetUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAdminServer).SetUserRole(ctx, req.(*SetUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatAdmin_DeleteUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatAdminServer).DeleteUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatAdmin_DeleteUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatAdminServer).DeleteUserRole(ctx, req.(*DeleteUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatAdmin_ServiceDesc is the grpc.ServiceDesc for ChatAdmin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatAdmin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.chat.admin.v1.ChatAdmin",
	HandlerType: (*ChatAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _ChatAdmin_Login_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _ChatAdmin_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _ChatAdmin_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _ChatAdmin_DeleteUser_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _ChatAdmin_UserInfo_Handler,
		},
		{
			MethodName: "UserList",
			Handler:    _ChatAdmin_UserList_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _ChatAdmin_CreateRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _ChatAdmin_UpdateRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _ChatAdmin_DeleteRole_Handler,
		},
		{
			MethodName: "RoleInfo",
			Handler:    _ChatAdmin_RoleInfo_Handler,
		},
		{
			MethodName: "RoleList",
			Handler:    _ChatAdmin_RoleList_Handler,
		},
		{
			MethodName: "CreatePermission",
			Handler:    _ChatAdmin_CreatePermission_Handler,
		},
		{
			MethodName: "UpdatePermission",
			Handler:    _ChatAdmin_UpdatePermission_Handler,
		},
		{
			MethodName: "DeletePermission",
			Handler:    _ChatAdmin_DeletePermission_Handler,
		},
		{
			MethodName: "PermissionInfo",
			Handler:    _ChatAdmin_PermissionInfo_Handler,
		},
		{
			MethodName: "PermissionList",
			Handler:    _ChatAdmin_PermissionList_Handler,
		},
		{
			MethodName: "SetUserRole",
			Handler:    _ChatAdmin_SetUserRole_Handler,
		},
		{
			MethodName: "DeleteUserRole",
			Handler:    _ChatAdmin_DeleteUserRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/admin.proto",
}
