package biz_test

import (
	"context"
	v1 "github.com/869413421/micro-chat/api/user/service/v1"
	"github.com/869413421/micro-chat/app/user/service/internal/biz"
	"testing"

	"github.com/869413421/micro-chat/app/user/service/internal/mocks/mrepo"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

// TestBizPermissionUsecaseCreate 测试创建角色
func TestBizPermissionUsecaseCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mPermissionRepo := mrepo.NewMockPermissionRepo(ctrl)
	permissionCase := biz.NewPermissionUsecase(mPermissionRepo, nil)

	Convey("Testing Biz PermissionUsecase Create Method", t, func() {
		permissionInfo := &biz.Permission{
			ID:        1,
			Name:      "test",
			Memo:      "test",
			Path:      "test",
			Method:    "GET",
			ParentID:  1,
			ParentIDS: "1",
		}
		createRequest := &v1.CreatePermissionRequest{
			Name:     "test",
			Memo:     "test",
			Path:     "test",
			Method:   "test",
			ParentId: 1,
		}
		mPermissionRepo.EXPECT().Create(ctx, gomock.Any()).Return(permissionInfo, nil)
		result, err := permissionCase.Create(ctx, createRequest)
		So(err, ShouldBeNil)
		So(result.ID, ShouldEqual, 1)
		So(result.Name, ShouldEqual, "test")
		So(result.Memo, ShouldEqual, "test")
		So(result.Path, ShouldEqual, "test")
		So(result.Method, ShouldEqual, "GET")
		So(result.ParentID, ShouldEqual, 1)
		So(result.ParentIDS, ShouldEqual, "1")
	})
}

// TestBizPermissionUsecaseUpdate 测试更新角色
func TestBizPermissionUsecaseUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mPermissionRepo := mrepo.NewMockPermissionRepo(ctrl)
	permissionCase := biz.NewPermissionUsecase(mPermissionRepo, nil)

	Convey("Testing Biz PermissionUsecase Update Method", t, func() {
		permissionInfo := &biz.Permission{
			ID:        1,
			Name:      "test",
			Memo:      "test",
			Path:      "test",
			Method:    "GET",
			ParentID:  1,
			ParentIDS: "1",
		}

		updateRequest := &v1.UpdatePermissionRequest{
			Id:       1,
			Name:     "test",
			Memo:     "test",
			Path:     "test",
			Method:   "test",
			ParentId: 1,
		}
		mPermissionRepo.EXPECT().Update(ctx, gomock.Any()).Return(permissionInfo, nil)
		result, err := permissionCase.Update(ctx, updateRequest)
		So(err, ShouldBeNil)
		So(result.ID, ShouldEqual, 1)
		So(result.Name, ShouldEqual, "test")
		So(result.Memo, ShouldEqual, "test")
		So(result.Path, ShouldEqual, "test")
		So(result.Method, ShouldEqual, "GET")
		So(result.ParentID, ShouldEqual, 1)
		So(result.ParentIDS, ShouldEqual, "1")
	})
}

// TestBizPermissionUsecaseDelete 测试删除角色
func TestBizPermissionUsecaseDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mPermissionRepo := mrepo.NewMockPermissionRepo(ctrl)
	permissionCase := biz.NewPermissionUsecase(mPermissionRepo, nil)

	Convey("Testing Biz PermissionUsecase Delete Method", t, func() {
		permissionInfo := &biz.Permission{
			ID:        1,
			Name:      "test",
			Memo:      "test",
			Path:      "test",
			Method:    "GET",
			ParentID:  1,
			ParentIDS: "1",
		}
		mPermissionRepo.EXPECT().Delete(ctx, gomock.Any()).Return(permissionInfo, nil)
		result, err := permissionCase.Delete(ctx, 1)
		So(err, ShouldBeNil)
		So(result.ID, ShouldEqual, 1)
		So(result.Name, ShouldEqual, "test")
		So(result.Memo, ShouldEqual, "test")
		So(result.Path, ShouldEqual, "test")
		So(result.Method, ShouldEqual, "GET")
		So(result.ParentID, ShouldEqual, 1)
		So(result.ParentIDS, ShouldEqual, "1")
	})
}

// TestBizPermissionUsecaseGet 测试获取角色
func TestBizPermissionUsecaseGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mPermissionRepo := mrepo.NewMockPermissionRepo(ctrl)
	permissionCase := biz.NewPermissionUsecase(mPermissionRepo, nil)

	Convey("Testing Biz PermissionUsecase Get Method", t, func() {
		permissionInfo := &biz.Permission{
			ID:        1,
			Name:      "test",
			Memo:      "test",
			Path:      "test",
			Method:    "GET",
			ParentID:  1,
			ParentIDS: "1",
		}
		mPermissionRepo.EXPECT().Get(ctx, gomock.Any()).Return(permissionInfo, nil)
		req := &v1.GetPermissionRequest{Id: 1}
		result, err := permissionCase.Get(ctx, req)
		So(err, ShouldBeNil)
		So(result.ID, ShouldEqual, 1)
		So(result.Name, ShouldEqual, "test")
		So(result.Memo, ShouldEqual, "test")
		So(result.Path, ShouldEqual, "test")
		So(result.Method, ShouldEqual, "GET")
		So(result.ParentID, ShouldEqual, 1)
		So(result.ParentIDS, ShouldEqual, "1")
	})
}

// TestBizPermissionUsecaseList 测试获取角色列表
func TestBizPermissionUsecaseList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mPermissionRepo := mrepo.NewMockPermissionRepo(ctrl)
	permissionCase := biz.NewPermissionUsecase(mPermissionRepo, nil)

	Convey("Testing Biz PermissionUsecase List Method", t, func() {
		permissionInfo := &biz.Permission{
			ID:        1,
			Name:      "test",
			Memo:      "test",
			Path:      "test",
			Method:    "GET",
			ParentID:  1,
			ParentIDS: "1",
		}
		Permissions := []*biz.Permission{permissionInfo}
		mPermissionRepo.EXPECT().List(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(Permissions, int64(1), nil)
		req := &v1.ListPermissionRequest{
			Page:     0,
			PageSize: 10,
		}
		result, total, err := permissionCase.List(ctx, req)
		permission := result[0]
		So(err, ShouldBeNil)
		So(len(result), ShouldEqual, 1)
		So(total, ShouldEqual, 1)
		So(err, ShouldBeNil)
		So(permission.ID, ShouldEqual, 1)
		So(permission.Name, ShouldEqual, "test")
		So(permission.Memo, ShouldEqual, "test")
		So(permission.Path, ShouldEqual, "test")
		So(permission.Method, ShouldEqual, "GET")
		So(permission.ParentID, ShouldEqual, 1)
		So(permission.ParentIDS, ShouldEqual, "1")
	})
}
