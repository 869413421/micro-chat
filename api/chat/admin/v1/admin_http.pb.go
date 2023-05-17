// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.18.0
// source: v1/admin.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationChatAdminCreatePermission = "/api.chat.admin.v1.ChatAdmin/CreatePermission"
const OperationChatAdminCreateRole = "/api.chat.admin.v1.ChatAdmin/CreateRole"
const OperationChatAdminCreateUser = "/api.chat.admin.v1.ChatAdmin/CreateUser"
const OperationChatAdminDeletePermission = "/api.chat.admin.v1.ChatAdmin/DeletePermission"
const OperationChatAdminDeleteRole = "/api.chat.admin.v1.ChatAdmin/DeleteRole"
const OperationChatAdminDeleteUser = "/api.chat.admin.v1.ChatAdmin/DeleteUser"
const OperationChatAdminLogin = "/api.chat.admin.v1.ChatAdmin/Login"
const OperationChatAdminPermissionInfo = "/api.chat.admin.v1.ChatAdmin/PermissionInfo"
const OperationChatAdminPermissionList = "/api.chat.admin.v1.ChatAdmin/PermissionList"
const OperationChatAdminRoleInfo = "/api.chat.admin.v1.ChatAdmin/RoleInfo"
const OperationChatAdminRoleList = "/api.chat.admin.v1.ChatAdmin/RoleList"
const OperationChatAdminUpdatePermission = "/api.chat.admin.v1.ChatAdmin/UpdatePermission"
const OperationChatAdminUpdateRole = "/api.chat.admin.v1.ChatAdmin/UpdateRole"
const OperationChatAdminUpdateUser = "/api.chat.admin.v1.ChatAdmin/UpdateUser"
const OperationChatAdminUserInfo = "/api.chat.admin.v1.ChatAdmin/UserInfo"
const OperationChatAdminUserList = "/api.chat.admin.v1.ChatAdmin/UserList"

type ChatAdminHTTPServer interface {
	CreatePermission(context.Context, *CreatePermissionRequest) (*PermissionInfoResponse, error)
	CreateRole(context.Context, *CreateRoleRequest) (*RoleInfoResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*UserInfoResponse, error)
	DeletePermission(context.Context, *DeletePermissionRequest) (*PermissionInfoResponse, error)
	DeleteRole(context.Context, *DeleteRoleRequest) (*RoleInfoResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*UserInfoResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	PermissionInfo(context.Context, *PermissionInfoRequest) (*PermissionInfoResponse, error)
	PermissionList(context.Context, *PermissionListRequest) (*PermissionListResponse, error)
	RoleInfo(context.Context, *RoleInfoRequest) (*RoleInfoResponse, error)
	RoleList(context.Context, *RoleListRequest) (*RoleListResponse, error)
	UpdatePermission(context.Context, *UpdatePermissionRequest) (*PermissionInfoResponse, error)
	UpdateRole(context.Context, *UpdateRoleRequest) (*RoleInfoResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UserInfoResponse, error)
	UserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error)
	UserList(context.Context, *UserListRequest) (*UserListResponse, error)
}

func RegisterChatAdminHTTPServer(s *http.Server, srv ChatAdminHTTPServer) {
	r := s.Route("/")
	r.POST("/admin/v1/login", _ChatAdmin_Login0_HTTP_Handler(srv))
	r.POST("/admin/v1/user", _ChatAdmin_CreateUser0_HTTP_Handler(srv))
	r.PUT("/admin/v1/user/{id}", _ChatAdmin_UpdateUser0_HTTP_Handler(srv))
	r.DELETE("/admin/v1/user/{id}", _ChatAdmin_DeleteUser0_HTTP_Handler(srv))
	r.GET("/admin/v1/user/{id}", _ChatAdmin_UserInfo0_HTTP_Handler(srv))
	r.GET("/admin/v1/user", _ChatAdmin_UserList0_HTTP_Handler(srv))
	r.POST("/admin/v1/role", _ChatAdmin_CreateRole0_HTTP_Handler(srv))
	r.PUT("/admin/v1/role/{id}", _ChatAdmin_UpdateRole0_HTTP_Handler(srv))
	r.DELETE("/admin/v1/role/{id}", _ChatAdmin_DeleteRole0_HTTP_Handler(srv))
	r.GET("/admin/v1/role/{id}", _ChatAdmin_RoleInfo0_HTTP_Handler(srv))
	r.GET("/admin/v1/role", _ChatAdmin_RoleList0_HTTP_Handler(srv))
	r.POST("/admin/v1/permission", _ChatAdmin_CreatePermission0_HTTP_Handler(srv))
	r.PUT("/admin/v1/permission/{id}", _ChatAdmin_UpdatePermission0_HTTP_Handler(srv))
	r.DELETE("/admin/v1/permission/{id}", _ChatAdmin_DeletePermission0_HTTP_Handler(srv))
	r.GET("/admin/v1/permission/{id}", _ChatAdmin_PermissionInfo0_HTTP_Handler(srv))
	r.GET("/admin/v1/permission", _ChatAdmin_PermissionList0_HTTP_Handler(srv))
}

func _ChatAdmin_Login0_HTTP_Handler(srv ChatAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatAdminLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginResponse)
		return ctx.Result(200, reply)
	}
}

func _ChatAdmin_CreateUser0_HTTP_Handler(srv ChatAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatAdminCreateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*CreateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoResponse)
		return ctx.Result(200, reply)
	}
}

func _ChatAdmin_UpdateUser0_HTTP_Handler(srv ChatAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatAdminUpdateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoResponse)
		return ctx.Result(200, reply)
	}
}

func _ChatAdmin_DeleteUser0_HTTP_Handler(srv ChatAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatAdminDeleteUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUser(ctx, req.(*DeleteUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoResponse)
		return ctx.Result(200, reply)
	}
}

func _ChatAdmin_UserInfo0_HTTP_Handler(srv ChatAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatAdminUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserInfo(ctx, req.(*UserInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserInfoResponse)
		return ctx.Result(200, reply)
	}
}

func _ChatAdmin_UserList0_HTTP_Handler(srv ChatAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatAdminUserList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserList(ctx, req.(*UserListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserListResponse)
		return ctx.Result(200, reply)
	}
}

func _ChatAdmin_CreateRole0_HTTP_Handler(srv ChatAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateRoleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatAdminCreateRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateRole(ctx, req.(*CreateRoleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoleInfoResponse)
		return ctx.Result(200, reply)
	}
}

func _ChatAdmin_UpdateRole0_HTTP_Handler(srv ChatAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateRoleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatAdminUpdateRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateRole(ctx, req.(*UpdateRoleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoleInfoResponse)
		return ctx.Result(200, reply)
	}
}

func _ChatAdmin_DeleteRole0_HTTP_Handler(srv ChatAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteRoleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatAdminDeleteRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteRole(ctx, req.(*DeleteRoleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoleInfoResponse)
		return ctx.Result(200, reply)
	}
}

func _ChatAdmin_RoleInfo0_HTTP_Handler(srv ChatAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RoleInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatAdminRoleInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoleInfo(ctx, req.(*RoleInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoleInfoResponse)
		return ctx.Result(200, reply)
	}
}

func _ChatAdmin_RoleList0_HTTP_Handler(srv ChatAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RoleListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatAdminRoleList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoleList(ctx, req.(*RoleListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoleListResponse)
		return ctx.Result(200, reply)
	}
}

func _ChatAdmin_CreatePermission0_HTTP_Handler(srv ChatAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreatePermissionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatAdminCreatePermission)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePermission(ctx, req.(*CreatePermissionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PermissionInfoResponse)
		return ctx.Result(200, reply)
	}
}

func _ChatAdmin_UpdatePermission0_HTTP_Handler(srv ChatAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePermissionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatAdminUpdatePermission)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePermission(ctx, req.(*UpdatePermissionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PermissionInfoResponse)
		return ctx.Result(200, reply)
	}
}

func _ChatAdmin_DeletePermission0_HTTP_Handler(srv ChatAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeletePermissionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatAdminDeletePermission)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeletePermission(ctx, req.(*DeletePermissionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PermissionInfoResponse)
		return ctx.Result(200, reply)
	}
}

func _ChatAdmin_PermissionInfo0_HTTP_Handler(srv ChatAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PermissionInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatAdminPermissionInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PermissionInfo(ctx, req.(*PermissionInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PermissionInfoResponse)
		return ctx.Result(200, reply)
	}
}

func _ChatAdmin_PermissionList0_HTTP_Handler(srv ChatAdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PermissionListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationChatAdminPermissionList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PermissionList(ctx, req.(*PermissionListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PermissionListResponse)
		return ctx.Result(200, reply)
	}
}

type ChatAdminHTTPClient interface {
	CreatePermission(ctx context.Context, req *CreatePermissionRequest, opts ...http.CallOption) (rsp *PermissionInfoResponse, err error)
	CreateRole(ctx context.Context, req *CreateRoleRequest, opts ...http.CallOption) (rsp *RoleInfoResponse, err error)
	CreateUser(ctx context.Context, req *CreateUserRequest, opts ...http.CallOption) (rsp *UserInfoResponse, err error)
	DeletePermission(ctx context.Context, req *DeletePermissionRequest, opts ...http.CallOption) (rsp *PermissionInfoResponse, err error)
	DeleteRole(ctx context.Context, req *DeleteRoleRequest, opts ...http.CallOption) (rsp *RoleInfoResponse, err error)
	DeleteUser(ctx context.Context, req *DeleteUserRequest, opts ...http.CallOption) (rsp *UserInfoResponse, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginResponse, err error)
	PermissionInfo(ctx context.Context, req *PermissionInfoRequest, opts ...http.CallOption) (rsp *PermissionInfoResponse, err error)
	PermissionList(ctx context.Context, req *PermissionListRequest, opts ...http.CallOption) (rsp *PermissionListResponse, err error)
	RoleInfo(ctx context.Context, req *RoleInfoRequest, opts ...http.CallOption) (rsp *RoleInfoResponse, err error)
	RoleList(ctx context.Context, req *RoleListRequest, opts ...http.CallOption) (rsp *RoleListResponse, err error)
	UpdatePermission(ctx context.Context, req *UpdatePermissionRequest, opts ...http.CallOption) (rsp *PermissionInfoResponse, err error)
	UpdateRole(ctx context.Context, req *UpdateRoleRequest, opts ...http.CallOption) (rsp *RoleInfoResponse, err error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest, opts ...http.CallOption) (rsp *UserInfoResponse, err error)
	UserInfo(ctx context.Context, req *UserInfoRequest, opts ...http.CallOption) (rsp *UserInfoResponse, err error)
	UserList(ctx context.Context, req *UserListRequest, opts ...http.CallOption) (rsp *UserListResponse, err error)
}

type ChatAdminHTTPClientImpl struct {
	cc *http.Client
}

func NewChatAdminHTTPClient(client *http.Client) ChatAdminHTTPClient {
	return &ChatAdminHTTPClientImpl{client}
}

func (c *ChatAdminHTTPClientImpl) CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...http.CallOption) (*PermissionInfoResponse, error) {
	var out PermissionInfoResponse
	pattern := "/admin/v1/permission"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChatAdminCreatePermission))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatAdminHTTPClientImpl) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...http.CallOption) (*RoleInfoResponse, error) {
	var out RoleInfoResponse
	pattern := "/admin/v1/role"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChatAdminCreateRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatAdminHTTPClientImpl) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...http.CallOption) (*UserInfoResponse, error) {
	var out UserInfoResponse
	pattern := "/admin/v1/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChatAdminCreateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatAdminHTTPClientImpl) DeletePermission(ctx context.Context, in *DeletePermissionRequest, opts ...http.CallOption) (*PermissionInfoResponse, error) {
	var out PermissionInfoResponse
	pattern := "/admin/v1/permission/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationChatAdminDeletePermission))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatAdminHTTPClientImpl) DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...http.CallOption) (*RoleInfoResponse, error) {
	var out RoleInfoResponse
	pattern := "/admin/v1/role/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationChatAdminDeleteRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatAdminHTTPClientImpl) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...http.CallOption) (*UserInfoResponse, error) {
	var out UserInfoResponse
	pattern := "/admin/v1/user/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationChatAdminDeleteUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatAdminHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginResponse, error) {
	var out LoginResponse
	pattern := "/admin/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChatAdminLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatAdminHTTPClientImpl) PermissionInfo(ctx context.Context, in *PermissionInfoRequest, opts ...http.CallOption) (*PermissionInfoResponse, error) {
	var out PermissionInfoResponse
	pattern := "/admin/v1/permission/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationChatAdminPermissionInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatAdminHTTPClientImpl) PermissionList(ctx context.Context, in *PermissionListRequest, opts ...http.CallOption) (*PermissionListResponse, error) {
	var out PermissionListResponse
	pattern := "/admin/v1/permission"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationChatAdminPermissionList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatAdminHTTPClientImpl) RoleInfo(ctx context.Context, in *RoleInfoRequest, opts ...http.CallOption) (*RoleInfoResponse, error) {
	var out RoleInfoResponse
	pattern := "/admin/v1/role/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationChatAdminRoleInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatAdminHTTPClientImpl) RoleList(ctx context.Context, in *RoleListRequest, opts ...http.CallOption) (*RoleListResponse, error) {
	var out RoleListResponse
	pattern := "/admin/v1/role"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationChatAdminRoleList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatAdminHTTPClientImpl) UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...http.CallOption) (*PermissionInfoResponse, error) {
	var out PermissionInfoResponse
	pattern := "/admin/v1/permission/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChatAdminUpdatePermission))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatAdminHTTPClientImpl) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...http.CallOption) (*RoleInfoResponse, error) {
	var out RoleInfoResponse
	pattern := "/admin/v1/role/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChatAdminUpdateRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatAdminHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...http.CallOption) (*UserInfoResponse, error) {
	var out UserInfoResponse
	pattern := "/admin/v1/user/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationChatAdminUpdateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatAdminHTTPClientImpl) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...http.CallOption) (*UserInfoResponse, error) {
	var out UserInfoResponse
	pattern := "/admin/v1/user/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationChatAdminUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ChatAdminHTTPClientImpl) UserList(ctx context.Context, in *UserListRequest, opts ...http.CallOption) (*UserListResponse, error) {
	var out UserListResponse
	pattern := "/admin/v1/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationChatAdminUserList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
