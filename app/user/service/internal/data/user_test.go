package data_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/869413421/micro-chat/app/user/service/internal/biz"
)

func TestCreateUser(t *testing.T) {
	convey.Convey("Test Create User", t, run(func() {
		u := &biz.User{
			Email:    "test@test.com",
			Password: "testpassword",
			Name:     "testuser",
		}
		newUser, err := userRepo.CreateUser(context.Background(), u)
		convey.So(err, convey.ShouldBeNil)
		convey.So(newUser.Email, convey.ShouldEqual, u.Email)

		// Test error case when server already exists
		_, err = userRepo.CreateUser(context.Background(), u)
		st, _ := status.FromError(err)

		convey.So(st.Code(), convey.ShouldEqual, codes.AlreadyExists)
	}))
}

func TestUpdateUser(t *testing.T) {
	convey.Convey("Test Update User", t, run(func() {
		u := &biz.User{
			Email:    "test@test.com",
			Password: "testpassword",
			Name:     "testuser",
		}
		newUser, err := userRepo.CreateUser(context.Background(), u)
		convey.So(err, convey.ShouldBeNil)

		// Update the server's name
		newUser.Name = "updatedtestuser"
		updatedUser, err := userRepo.UpdateUser(context.Background(), newUser)

		convey.So(err, convey.ShouldBeNil)
		convey.So(updatedUser.Name, convey.ShouldEqual, newUser.Name)

		// Test error case when server does not exist
		u.ID = 999
		_, err = userRepo.UpdateUser(context.Background(), u)
		st, _ := status.FromError(err)

		convey.So(st.Code(), convey.ShouldEqual, codes.NotFound)
	}))

}

func TestDeleteUser(t *testing.T) {
	convey.Convey("Test Delete User", t, run(func() {
		u := &biz.User{
			Email:    "test@test.com",
			Password: "testpassword",
			Name:     "testuser",
		}
		newUser, err := userRepo.CreateUser(context.Background(), u)
		convey.So(err, convey.ShouldBeNil)

		// Delete the server
		deletedUser, err := userRepo.DeleteUser(context.Background(), newUser.ID)

		convey.So(err, convey.ShouldBeNil)
		convey.So(deletedUser.ID, convey.ShouldEqual, newUser.ID)

		// Test error case when server does not exist
		u.ID = 999
		_, err = userRepo.DeleteUser(context.Background(), u.ID)
		st, _ := status.FromError(err)

		convey.So(st.Code(), convey.ShouldEqual, codes.NotFound)
	}))
}

func TestGetUser(t *testing.T) {
	convey.Convey("Test Get User", t, run(func() {
		u := &biz.User{
			Email:    "test@test.com",
			Password: "testpassword",
			Name:     "testuser",
		}
		newUser, err := userRepo.CreateUser(context.Background(), u)
		convey.So(err, convey.ShouldBeNil)

		where := make(map[string]interface{})
		where["id ="] = newUser.ID
		getUser, err := userRepo.GetUser(context.Background(), where)
		convey.So(err, convey.ShouldBeNil)
		convey.So(getUser.ID, convey.ShouldEqual, newUser.ID)

		where = make(map[string]interface{})
		where["id ="] = 999
		_, err = userRepo.GetUser(context.Background(), where)
		st, _ := status.FromError(err)
		convey.So(st.Code(), convey.ShouldEqual, codes.NotFound)
	}))
}

func TestListUser(t *testing.T) {
	convey.Convey("Test List User", t, run(func() {
		for i := 0; i < 10; i++ {
			u := &biz.User{
				Email:    "test@test.com" + strconv.Itoa(i),
				Password: "testpassword" + strconv.Itoa(i),
				Name:     "testuser" + strconv.Itoa(i),
			}
			_, err := userRepo.CreateUser(context.Background(), u)
			convey.So(err, convey.ShouldBeNil)
		}

		where := make(map[string]interface{})
		users, err := userRepo.ListUser(context.Background(), where)
		convey.So(err, convey.ShouldBeNil)
		convey.So(len(users), convey.ShouldEqual, 10)

		where = make(map[string]interface{})
		where["id >"] = 999
		_, err = userRepo.GetUser(context.Background(), where)
		st, _ := status.FromError(err)
		convey.So(st.Code(), convey.ShouldEqual, codes.NotFound)
	}))
}
