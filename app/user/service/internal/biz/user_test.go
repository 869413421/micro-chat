package biz_test

import (
	"context"
	"fmt"
	"github.com/869413421/micro-chat/app/user/service/internal/biz"
	"strconv"
	"testing"

	"github.com/869413421/micro-chat/app/user/service/internal/mocks/mrepo"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBizUserUsecaseCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mUserRepo := mrepo.NewMockUserRepo(ctrl)
	userCase := biz.NewUserUsecase(mUserRepo, nil, nil)

	Convey("Testing Biz UserUsecase Create Method", t, func() {
		userInfo := &biz.User{
			ID:       1,
			Name:     "清水泥沙",
			Email:    "13528685024@163",
			Password: "server",
		}
		mUserRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return(userInfo, nil)
		result, err := userCase.Create(ctx, userInfo)
		So(err, ShouldBeNil)
		So(result.ID, ShouldEqual, 1)
		So(result.Name, ShouldEqual, "清水泥沙")
		So(result.Email, ShouldEqual, "13528685024@163")
		So(result.Password, ShouldEqual, "server")
	})
}

func TestBizUserUsecaseUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mUserRepo := mrepo.NewMockUserRepo(ctrl)
	userCase := biz.NewUserUsecase(mUserRepo, nil, nil)

	Convey("Testing Biz UserUsecase Update Method", t, func() {
		userInfo := &biz.User{
			ID:       1,
			Name:     "清水泥沙",
			Email:    "13528685024@163",
			Password: "server",
		}
		mUserRepo.EXPECT().UpdateUser(ctx, gomock.Any()).Return(userInfo, nil)
		result, err := userCase.Update(ctx, userInfo)
		So(err, ShouldBeNil)
		So(result.ID, ShouldEqual, 1)
		So(result.Name, ShouldEqual, "清水泥沙")
		So(result.Email, ShouldEqual, "13528685024@163")
		So(result.Password, ShouldEqual, "server")
	})
}

func TestBizUserUsecaseDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mUserRepo := mrepo.NewMockUserRepo(ctrl)
	userCase := biz.NewUserUsecase(mUserRepo, nil, nil)

	Convey("Testing Biz UserUsecase Delete Method", t, func() {
		userInfo := &biz.User{
			ID:       1,
			Name:     "清水泥沙",
			Email:    "13528685024@163",
			Password: "server",
		}
		mUserRepo.EXPECT().DeleteUser(ctx, gomock.Any()).Return(userInfo, nil)
		result, err := userCase.Delete(ctx, userInfo.ID)
		So(err, ShouldBeNil)
		So(result.ID, ShouldEqual, 1)
		So(result.Name, ShouldEqual, "清水泥沙")
		So(result.Email, ShouldEqual, "13528685024@163")
		So(result.Password, ShouldEqual, "server")
	})
}

func TestBizUserUsecaseGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mUserRepo := mrepo.NewMockUserRepo(ctrl)
	userCase := biz.NewUserUsecase(mUserRepo, nil, nil)

	Convey("Testing Biz UserUsecase Get Method", t, func() {
		userInfo := &biz.User{
			ID:       1,
			Name:     "清水泥沙",
			Email:    "13528685024@163",
			Password: "server",
		}
		mUserRepo.EXPECT().GetUser(ctx, gomock.Any()).Return(userInfo, nil)
		var where = map[string]interface{}{"id": 1}
		result, err := userCase.Get(ctx, where)
		So(err, ShouldBeNil)
		So(result.ID, ShouldEqual, 1)
		So(result.Name, ShouldEqual, "清水泥沙")
		So(result.Email, ShouldEqual, "13528685024@163")
		So(result.Password, ShouldEqual, "server")
	})
}

func TestBizUserUsecaseList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mUserRepo := mrepo.NewMockUserRepo(ctrl)
	userCase := biz.NewUserUsecase(mUserRepo, nil, nil)

	Convey("Testing Biz UserUsecase List Method", t, func() {
		var list []*biz.User
		for i := 0; i < 10; i++ {
			list = append(list, &biz.User{
				ID:       uint64(i),
				Name:     "清水泥沙",
				Email:    fmt.Sprintf("1352868502%s@163", strconv.Itoa(i)),
				Password: "server",
			})
		}
		mUserRepo.EXPECT().ListUser(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return(list, int64(10), nil)
		var where = map[string]interface{}{"id": 1}
		_, total, err := userCase.List(ctx, where, 1, 10)
		So(err, ShouldBeNil)
		So(total, ShouldEqual, 10)
	})
}
