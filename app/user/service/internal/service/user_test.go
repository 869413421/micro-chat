package service_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/869413421/micro-chat/api/user/service/v1"
	"testing"

	"github.com/869413421/micro-chat/app/user/service/internal/biz"
	mock_biz "github.com/869413421/micro-chat/app/user/service/internal/mocks/mbiz"
	"github.com/869413421/micro-chat/app/user/service/internal/service"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

// TestUserService_CreateUser 测试创建用户
func TestUserService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUser := &biz.User{
		Name:     "Test User",
		Password: "123456",
		Email:    "test@example.com",
	}

	mockUsecase := mock_biz.NewMockUserUsecase(ctrl)
	mockUsecase.EXPECT().Create(gomock.Any(), gomock.Any()).Return(mockUser, nil)

	svc := service.NewUserService(mockUsecase, nil)

	Convey("Test Service CreateUser", t, func() {
		resp, err := svc.CreateUser(context.Background(), &v1.CreateUserRequest{
			Name:     mockUser.Name,
			Password: mockUser.Password,
			Email:    mockUser.Email,
		})

		So(err, ShouldBeNil)
		So(resp, ShouldNotBeNil)
		So(resp.Name, ShouldEqual, mockUser.Name)
		So(resp.Password, ShouldEqual, mockUser.Password)
		So(resp.Email, ShouldEqual, mockUser.Email)
	})
}

// TestUserService_CreateUser_Error 测试创建用户错误
func TestUserService_CreateUser_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_biz.NewMockUserUsecase(ctrl)
	mockUsecase.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("Create server error"))

	svc := service.NewUserService(mockUsecase, nil)

	Convey("Test Service CreateUser Error", t, func() {
		resp, err := svc.CreateUser(context.Background(), &v1.CreateUserRequest{
			Name:     "Test User",
			Password: "123456",
			Email:    "test@example.com",
		})

		So(resp, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})
}

// TestUserService_UpdateUser 测试更新用户
func TestUserService_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUser := &biz.User{
		ID:       1,
		Name:     "Test User",
		Password: "123456",
		Email:    "test@example.com",
	}

	mockUsecase := mock_biz.NewMockUserUsecase(ctrl)
	mockUsecase.EXPECT().Update(gomock.Any(), gomock.Any()).Return(mockUser, nil)

	svc := service.NewUserService(mockUsecase, nil)

	Convey("Test Service UpdateUser Error", t, func() {
		resp, err := svc.UpdateUser(context.Background(), &v1.UpdateUserRequest{
			Id:       mockUser.ID,
			Name:     mockUser.Name,
			Password: mockUser.Password,
			Email:    mockUser.Email,
		})

		So(err, ShouldBeNil)
		So(resp, ShouldNotBeNil)
		So(resp.Name, ShouldEqual, mockUser.Name)
		So(resp.Password, ShouldEqual, mockUser.Password)
		So(resp.Email, ShouldEqual, mockUser.Email)
	})
}

// TestUserService_UpdateUser_Error 测试更新用户错误
func TestUserService_UpdateUser_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUser := &biz.User{
		ID:       1,
		Name:     "Test User",
		Password: "123456",
		Email:    "test@example.com",
	}

	mockUsecase := mock_biz.NewMockUserUsecase(ctrl)
	mockUsecase.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil, errors.New("server update error"))

	svc := service.NewUserService(mockUsecase, nil)

	Convey("Test Service UpdateUser Error", t, func() {
		resp, err := svc.UpdateUser(context.Background(), &v1.UpdateUserRequest{
			Id:       mockUser.ID,
			Name:     mockUser.Name,
			Password: mockUser.Password,
			Email:    mockUser.Email,
		})

		So(resp, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})
}

// TestUserService_DeleteUser 测试删除用户
func TestUserService_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUser := &biz.User{
		ID:       1,
		Name:     "Test User",
		Password: "123456",
		Email:    "test@example.com",
	}
	mockUsecase := mock_biz.NewMockUserUsecase(ctrl)
	mockUsecase.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(mockUser, nil)

	svc := service.NewUserService(mockUsecase, nil)

	Convey("Test Service DeleteUser", t, func() {
		resp, err := svc.DeleteUser(context.Background(), &v1.DeleteUserRequest{
			Id: 1,
		})

		So(err, ShouldBeNil)
		So(resp, ShouldNotBeNil)
		So(resp.Name, ShouldEqual, mockUser.Name)
		So(resp.Password, ShouldEqual, mockUser.Password)
		So(resp.Email, ShouldEqual, mockUser.Email)
	})
}

// TestUserService_DeleteUser_Error 测试删除用户错误
func TestUserService_DeleteUser_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_biz.NewMockUserUsecase(ctrl)
	mockUsecase.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil, errors.New("server delete error"))

	svc := service.NewUserService(mockUsecase, nil)

	Convey("Test Service DeleteUser Error", t, func() {
		resp, err := svc.DeleteUser(context.Background(), &v1.DeleteUserRequest{
			Id: 1,
		})

		So(resp, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})
}

// TestUserService_GetUser 测试获取用户
func TestUserService_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUser := &biz.User{
		ID:       1,
		Name:     "Test User",
		Password: "123456",
		Email:    "test@example.com",
	}

	mockUsecase := mock_biz.NewMockUserUsecase(ctrl)
	mockUsecase.EXPECT().Get(gomock.Any(), gomock.Any()).Return(mockUser, nil)

	svc := service.NewUserService(mockUsecase, nil)

	Convey("Test Service GetUser", t, func() {
		resp, err := svc.GetUser(context.Background(), &v1.GetUserRequest{
			Email: "test@example.com",
		})

		So(err, ShouldBeNil)
		So(resp, ShouldNotBeNil)
		So(resp.Name, ShouldEqual, mockUser.Name)
		So(resp.Password, ShouldEqual, mockUser.Password)
		So(resp.Email, ShouldEqual, mockUser.Email)
	})
}

// TestUserService_GetUser_Error 测试获取用户错误
func TestUserService_GetUser_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_biz.NewMockUserUsecase(ctrl)
	mockUsecase.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, errors.New("server get error"))

	svc := service.NewUserService(mockUsecase, nil)

	Convey("Test Service GetUser Error", t, func() {
		resp, err := svc.GetUser(context.Background(), &v1.GetUserRequest{
			Email: "test@example.com",
		})

		So(resp, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})
}

// TestUserService_ListUser 测试获取用户列表
func TestUserService_ListUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var mockUsers []*biz.User
	for i := 0; i < 10; i++ {
		mockUsers = append(mockUsers, &biz.User{
			ID:    uint64(i),
			Name:  fmt.Sprintf("Test User %d", i),
			Email: fmt.Sprintf("test%d@example.com", i),
		})
	}

	mockUsecase := mock_biz.NewMockUserUsecase(ctrl)
	mockUsecase.EXPECT().ListUser(gomock.Any(), gomock.Any()).Return(mockUsers, nil)

	svc := service.NewUserService(mockUsecase, nil)
	Convey("Test Service GetUser Error", t, func() {
		resp, err := svc.ListUser(context.Background(), &v1.ListUserRequest{
			Offset: 0,
			Limit:  10,
		})

		So(err, ShouldBeNil)
		So(resp, ShouldNotBeNil)
		So(len(resp.Users), ShouldEqual, 10)
	})
}

// TestUserService_ListUser_Error 测试获取用户列表错误
func TestUserService_ListUser_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_biz.NewMockUserUsecase(ctrl)
	mockUsecase.EXPECT().ListUser(gomock.Any(), gomock.Any()).Return(nil, errors.New("not found server"))

	svc := service.NewUserService(mockUsecase, nil)
	Convey("Test Service GetUser Error", t, func() {
		resp, err := svc.ListUser(context.Background(), &v1.ListUserRequest{
			Offset: 0,
			Limit:  10,
		})

		So(err, ShouldNotBeNil)
		So(resp, ShouldBeNil)
	})
}
