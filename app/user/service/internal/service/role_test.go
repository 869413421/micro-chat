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

// TestUserService_CreateRole 测试创建角色
func TestUserService_CreateRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRole := &biz.Role{
		ID:   1,
		Name: "test",
		Memo: "test",
	}
	mockUsecase := mock_biz.NewMockRoleUsecase(ctrl)
	mockUsecase.EXPECT().Create(gomock.Any(), gomock.Any()).Return(mockRole, nil)

	svc := NewUserService(nil, mockUsecase, nil, nil)
	Convey("Test Service CreateRole", t, func() {
		resp, err := svc.CreateRole(context.Background(), &v1.CreateRoleRequest{
			Name: mockRole.Name,
			Memo: mockRole.Memo,
		})
		So(err, ShouldBeNil)
		So(resp.Id, ShouldEqual, mockRole.ID)
		So(resp.Name, ShouldEqual, mockRole.Name)
		So(resp.Memo, ShouldEqual, mockRole.Memo)
	})
}

// TestUserService_UpdateRole 测试更新角色
func TestUserService_UpdateRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRole := &biz.Role{
		ID:   1,
		Name: "test",
		Memo: "test",
	}
	mockUsecase := mock_biz.NewMockRoleUsecase(ctrl)
	mockUsecase.EXPECT().Update(gomock.Any(), gomock.Any()).Return(mockRole, nil)

	svc := NewUserService(nil, mockUsecase, nil, nil)
	Convey("Test Service UpdateRole", t, func() {
		resp, err := svc.UpdateRole(context.Background(), &v1.UpdateRoleRequest{
			Id:   1,
			Name: mockRole.Name,
			Memo: mockRole.Memo,
		})
		So(err, ShouldBeNil)
		So(resp.Id, ShouldEqual, mockRole.ID)
		So(resp.Name, ShouldEqual, mockRole.Name)
		So(resp.Memo, ShouldEqual, mockRole.Memo)
	})
}

// TestUserService_DeleteRole 测试删除角色
func TestUserService_DeleteRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	role := &biz.Role{
		ID:   1,
		Name: "test",
		Memo: "test",
	}
	mockUsecase := mock_biz.NewMockRoleUsecase(ctrl)
	mockUsecase.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(role, nil)

	svc := NewUserService(nil, mockUsecase, nil, nil)
	Convey("Test Service DeleteRole", t, func() {
		dRole, err := svc.DeleteRole(context.Background(), &v1.DeleteRoleRequest{
			Id: 1,
		})
		So(err, ShouldBeNil)
		So(dRole.Id, ShouldEqual, role.ID)
		So(dRole.Name, ShouldEqual, role.Name)
		So(dRole.Memo, ShouldEqual, role.Memo)
	})
}

// TestUserService_GetRole 测试获取角色
func TestUserService_GetRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	role := &biz.Role{
		ID:   1,
		Name: "test",
		Memo: "test",
	}
	mockUsecase := mock_biz.NewMockRoleUsecase(ctrl)
	mockUsecase.EXPECT().Get(gomock.Any(), gomock.Any()).Return(role, nil)

	svc := NewUserService(nil, mockUsecase, nil, nil)
	Convey("Test Service GetRole", t, func() {
		gRole, err := svc.GetRole(context.Background(), &v1.GetRoleRequest{
			Id: 1,
		})
		So(err, ShouldBeNil)
		So(gRole.Id, ShouldEqual, role.ID)
		So(gRole.Name, ShouldEqual, role.Name)
		So(gRole.Memo, ShouldEqual, role.Memo)
	})
}

// TestUserService_ListRole 测试获取角色列表
func TestUserService_ListRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	role := &biz.Role{
		ID:   1,
		Name: "test",
		Memo: "test",
	}
	mockUsecase := mock_biz.NewMockRoleUsecase(ctrl)
	mockUsecase.EXPECT().List(gomock.Any(), gomock.Any()).Return([]*biz.Role{role}, int64(1), nil)

	svc := NewUserService(nil, mockUsecase, nil, nil)
	Convey("Test Service ListRole", t, func() {
		response, err := svc.ListRole(context.Background(), &v1.ListRoleRequest{
			Page:     0,
			PageSize: 10,
		})
		So(err, ShouldBeNil)
		So(response.Total, ShouldEqual, 1)
		So(response.Roles[0].Id, ShouldEqual, role.ID)
		So(response.Roles[0].Name, ShouldEqual, role.Name)
		So(response.Roles[0].Memo, ShouldEqual, role.Memo)
	})
}
