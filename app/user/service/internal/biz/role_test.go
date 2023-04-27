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

// TestBizRoleUsecaseCreate 测试创建角色
func TestBizRoleUsecaseCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mRoleRepo := mrepo.NewMockRoleRepo(ctrl)
	roleCase := biz.NewRoleUsecase(mRoleRepo, nil)

	Convey("Testing Biz RoleUsecase Create Method", t, func() {
		roleInfo := &biz.Role{
			ID:   1,
			Name: "test",
			Memo: "test",
		}
		createRequest := &v1.CreateRoleRequest{
			Name: "test",
			Memo: "test",
		}
		mRoleRepo.EXPECT().Create(ctx, gomock.Any()).Return(roleInfo, nil)
		result, err := roleCase.Create(ctx, createRequest)
		So(err, ShouldBeNil)
		So(result.ID, ShouldEqual, 1)
		So(result.Name, ShouldEqual, "test")
		So(result.Memo, ShouldEqual, "test")
	})
}

// TestBizRoleUsecaseUpdate 测试更新角色
func TestBizRoleUsecaseUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mRoleRepo := mrepo.NewMockRoleRepo(ctrl)
	roleCase := biz.NewRoleUsecase(mRoleRepo, nil)

	Convey("Testing Biz RoleUsecase Update Method", t, func() {
		roleInfo := &biz.Role{
			ID:   1,
			Name: "test",
			Memo: "test",
		}
		updateRequest := &v1.UpdateRoleRequest{
			Id:   1,
			Name: "test",
			Memo: "test",
		}
		mRoleRepo.EXPECT().Update(ctx, gomock.Any()).Return(roleInfo, nil)
		result, err := roleCase.Update(ctx, updateRequest)
		So(err, ShouldBeNil)
		So(result.ID, ShouldEqual, 1)
		So(result.Name, ShouldEqual, "test")
		So(result.Memo, ShouldEqual, "test")
	})
}

// TestBizRoleUsecaseDelete 测试删除角色
func TestBizRoleUsecaseDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mRoleRepo := mrepo.NewMockRoleRepo(ctrl)
	roleCase := biz.NewRoleUsecase(mRoleRepo, nil)

	Convey("Testing Biz RoleUsecase Delete Method", t, func() {
		roleInfo := &biz.Role{
			ID:   1,
			Name: "test",
			Memo: "test",
		}
		mRoleRepo.EXPECT().Delete(ctx, gomock.Any()).Return(roleInfo, nil)
		result, err := roleCase.Delete(ctx, 1)
		So(err, ShouldBeNil)
		So(result.ID, ShouldEqual, 1)
		So(result.Name, ShouldEqual, "test")
		So(result.Memo, ShouldEqual, "test")
	})
}

// TestBizRoleUsecaseGet 测试获取角色
func TestBizRoleUsecaseGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mRoleRepo := mrepo.NewMockRoleRepo(ctrl)
	roleCase := biz.NewRoleUsecase(mRoleRepo, nil)

	Convey("Testing Biz RoleUsecase Get Method", t, func() {
		roleInfo := &biz.Role{
			ID:   1,
			Name: "test",
			Memo: "test",
		}
		mRoleRepo.EXPECT().Get(ctx, gomock.Any()).Return(roleInfo, nil)
		req := &v1.GetRoleRequest{Id: 1}
		result, err := roleCase.Get(ctx, req)
		So(err, ShouldBeNil)
		So(result.ID, ShouldEqual, 1)
		So(result.Name, ShouldEqual, "test")
		So(result.Memo, ShouldEqual, "test")
	})
}

// TestBizRoleUsecaseList 测试获取角色列表
func TestBizRoleUsecaseList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mRoleRepo := mrepo.NewMockRoleRepo(ctrl)
	roleCase := biz.NewRoleUsecase(mRoleRepo, nil)

	Convey("Testing Biz RoleUsecase List Method", t, func() {
		roleInfo := &biz.Role{
			ID:   1,
			Name: "test",
			Memo: "test",
		}
		roles := []*biz.Role{roleInfo}
		mRoleRepo.EXPECT().List(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(roles, int64(1), nil)
		req := &v1.ListRoleRequest{
			Offset: 0,
			Limit:  10,
		}
		result, total, err := roleCase.List(ctx, req)
		So(err, ShouldBeNil)
		So(len(result), ShouldEqual, 1)
		So(total, ShouldEqual, 1)
		So(result[0].ID, ShouldEqual, 1)
		So(result[0].Name, ShouldEqual, "test")
		So(result[0].Memo, ShouldEqual, "test")
	})
}
