package service

import (
	"context"
	"github.com/869413421/micro-chat/app/user/service/internal/biz"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"

	v1 "github.com/869413421/micro-chat/api/user/service/v1"
	mock_biz "github.com/869413421/micro-chat/app/user/service/internal/mocks/mbiz"
)

// TestUserService_CreatePermission 测试创建角色
func TestUserService_CreatePermission(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPermission := &biz.Permission{
		ID:     1,
		Name:   "test",
		Memo:   "test",
		Path:   "/test",
		Method: "GET",
	}
	mockUsecase := mock_biz.NewMockPermissionUsecase(ctrl)
	mockUsecase.EXPECT().Create(gomock.Any(), gomock.Any()).Return(mockPermission, nil)

	svc := NewUserService(nil, nil, mockUsecase, nil)
	Convey("Test Service CreatePermission", t, func() {
		resp, err := svc.CreatePermission(context.Background(), &v1.CreatePermissionRequest{
			Name: mockPermission.Name,
			Memo: mockPermission.Memo,
		})
		So(err, ShouldBeNil)
		So(resp.Id, ShouldEqual, mockPermission.ID)
		So(resp.Name, ShouldEqual, mockPermission.Name)
		So(resp.Memo, ShouldEqual, mockPermission.Memo)
		So(resp.Path, ShouldEqual, mockPermission.Path)
		So(resp.Method, ShouldEqual, mockPermission.Method)
	})
}

// TestUserService_UpdatePermission 测试更新角色
func TestUserService_UpdatePermission(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPermission := &biz.Permission{
		ID:     1,
		Name:   "test",
		Memo:   "test",
		Path:   "/test",
		Method: "GET",
	}
	mockUsecase := mock_biz.NewMockPermissionUsecase(ctrl)
	mockUsecase.EXPECT().Update(gomock.Any(), gomock.Any()).Return(mockPermission, nil)

	svc := NewUserService(nil, nil, mockUsecase, nil)
	Convey("Test Service UpdatePermission", t, func() {
		resp, err := svc.UpdatePermission(context.Background(), &v1.UpdatePermissionRequest{
			Id:   1,
			Name: mockPermission.Name,
			Memo: mockPermission.Memo,
		})
		So(err, ShouldBeNil)
		So(resp.Id, ShouldEqual, mockPermission.ID)
		So(resp.Name, ShouldEqual, mockPermission.Name)
		So(resp.Memo, ShouldEqual, mockPermission.Memo)
		So(resp.Path, ShouldEqual, mockPermission.Path)
		So(resp.Method, ShouldEqual, mockPermission.Method)
	})
}

// TestUserService_DeletePermission 测试删除角色
func TestUserService_DeletePermission(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	permission := &biz.Permission{
		ID:     1,
		Name:   "test",
		Memo:   "test",
		Path:   "/test",
		Method: "GET",
	}
	mockUsecase := mock_biz.NewMockPermissionUsecase(ctrl)
	mockUsecase.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(permission, nil)

	svc := NewUserService(nil, nil, mockUsecase, nil)
	Convey("Test Service DeletePermission", t, func() {
		dPermission, err := svc.DeletePermission(context.Background(), &v1.DeletePermissionRequest{
			Id: 1,
		})
		So(err, ShouldBeNil)
		So(dPermission.Id, ShouldEqual, permission.ID)
		So(dPermission.Name, ShouldEqual, permission.Name)
		So(dPermission.Memo, ShouldEqual, permission.Memo)
		So(dPermission.Path, ShouldEqual, permission.Path)
		So(dPermission.Method, ShouldEqual, permission.Method)
	})
}

// TestUserService_GetPermission 测试获取角色
func TestUserService_GetPermission(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	permission := &biz.Permission{
		ID:     1,
		Name:   "test",
		Memo:   "test",
		Path:   "/test",
		Method: "GET",
	}
	mockUsecase := mock_biz.NewMockPermissionUsecase(ctrl)
	mockUsecase.EXPECT().Get(gomock.Any(), gomock.Any()).Return(permission, nil)

	svc := NewUserService(nil, nil, mockUsecase, nil)
	Convey("Test Service GetPermission", t, func() {
		gPermission, err := svc.GetPermission(context.Background(), &v1.GetPermissionRequest{
			Id: 1,
		})
		So(err, ShouldBeNil)
		So(gPermission.Id, ShouldEqual, permission.ID)
		So(gPermission.Name, ShouldEqual, permission.Name)
		So(gPermission.Memo, ShouldEqual, permission.Memo)
		So(gPermission.Path, ShouldEqual, permission.Path)
		So(gPermission.Method, ShouldEqual, permission.Method)
	})
}

// TestUserService_ListPermission 测试获取角色列表
func TestUserService_ListPermission(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	permission := &biz.Permission{
		ID:     1,
		Name:   "test",
		Memo:   "test",
		Path:   "/test",
		Method: "GET",
	}
	mockUsecase := mock_biz.NewMockPermissionUsecase(ctrl)
	mockUsecase.EXPECT().List(gomock.Any(), gomock.Any()).Return([]*biz.Permission{permission}, int64(1), nil)

	svc := NewUserService(nil, nil, mockUsecase, nil)
	Convey("Test Service ListPermission", t, func() {
		response, err := svc.ListPermission(context.Background(), &v1.ListPermissionRequest{
			Page:     0,
			PageSize: 10,
		})
		So(err, ShouldBeNil)
		So(response.Total, ShouldEqual, 1)
		So(response.Permissions[0].Id, ShouldEqual, permission.ID)
		So(response.Permissions[0].Name, ShouldEqual, permission.Name)
		So(response.Permissions[0].Memo, ShouldEqual, permission.Memo)
		So(response.Permissions[0].Path, ShouldEqual, permission.Path)
		So(response.Permissions[0].Method, ShouldEqual, permission.Method)
	})
}
