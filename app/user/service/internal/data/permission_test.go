package data_test

import (
	"context"
	"github.com/869413421/micro-chat/app/user/service/internal/biz"
	"github.com/869413421/micro-chat/pkg/types"
	"strconv"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TestCreatePermission 测试创建权限
func TestCreatePermission(t *testing.T) {
	convey.Convey("Test Create Permission", t, run(func() {
		r := &biz.Permission{
			Name:     "test",
			Memo:     "test",
			Path:     "/test",
			Method:   "GET",
			ParentID: 0,
		}
		newPermission, err := permissionRepo.Create(context.Background(), r)
		convey.So(err, convey.ShouldBeNil)
		convey.So(newPermission.Name, convey.ShouldEqual, r.Name)
		convey.So(newPermission.ID, convey.ShouldNotEqual, 0)

		// Test error case when server already exists
		_, err = permissionRepo.Create(context.Background(), r)
		st, _ := status.FromError(err)
		convey.So(st.Code(), convey.ShouldEqual, codes.AlreadyExists)

		// 测试父级ID
		r2 := &biz.Permission{
			Name:     "test2",
			Memo:     "test2",
			Path:     "/test2",
			Method:   "GET",
			ParentID: newPermission.ID,
		}
		newPermission2, err := permissionRepo.Create(context.Background(), r2)
		convey.So(err, convey.ShouldBeNil)
		convey.So(newPermission2.Name, convey.ShouldEqual, r2.Name)
		convey.So(newPermission2.ParentID, convey.ShouldEqual, newPermission.ID)
		convey.So(newPermission2.ParentIDS, convey.ShouldEqual, types.UInt64ToString(newPermission.ID))
	}))
}

// TestUpdatePermission 测试更新权限
func TestUpdatePermission(t *testing.T) {
	convey.Convey("Test Update Permission", t, run(func() {
		r := &biz.Permission{
			Name:     "test",
			Memo:     "test",
			Path:     "/test",
			Method:   "GET",
			ParentID: 0,
		}
		newPermission, err := permissionRepo.Create(context.Background(), r)
		convey.So(err, convey.ShouldBeNil)
		convey.So(newPermission.Name, convey.ShouldEqual, r.Name)
		convey.So(newPermission.Memo, convey.ShouldEqual, r.Memo)
		convey.So(newPermission.Path, convey.ShouldEqual, r.Path)
		convey.So(newPermission.Method, convey.ShouldEqual, r.Method)
		convey.So(newPermission.ID, convey.ShouldNotEqual, 0)

		// 测试更新权限名称
		r.ID = newPermission.ID
		r.Name = "test2"
		newPermission, err = permissionRepo.Update(context.Background(), r)
		convey.So(err, convey.ShouldBeNil)
		convey.So(newPermission.Name, convey.ShouldEqual, "test2")
	}))
}

// TestUpdatePermission_Err 测试更新权限错误
func TestUpdatePermission_Err(t *testing.T) {
	convey.Convey("Test Update Permission", t, run(func() {
		r := &biz.Permission{
			Name:   "test",
			Memo:   "test",
			Path:   "/test",
			Method: "GET",
		}
		newPermission, err := permissionRepo.Create(context.Background(), r)
		convey.So(err, convey.ShouldBeNil)
		convey.So(newPermission.Name, convey.ShouldEqual, r.Name)
		convey.So(newPermission.Memo, convey.ShouldEqual, r.Memo)
		convey.So(newPermission.Path, convey.ShouldEqual, r.Path)
		convey.So(newPermission.Method, convey.ShouldEqual, r.Method)
		convey.So(newPermission.ID, convey.ShouldNotEqual, 0)

		r2 := &biz.Permission{
			Name:   "test2",
			Memo:   "test2",
			Path:   "/test2",
			Method: "GET",
		}
		newPermission2, err := permissionRepo.Create(context.Background(), r2)
		convey.So(err, convey.ShouldBeNil)
		convey.So(newPermission2.Name, convey.ShouldEqual, r2.Name)
		convey.So(newPermission2.Memo, convey.ShouldEqual, r2.Memo)
		convey.So(newPermission2.Path, convey.ShouldEqual, r2.Path)
		convey.So(newPermission2.Method, convey.ShouldEqual, r2.Method)
		convey.So(newPermission2.ID, convey.ShouldNotEqual, 0)

		// 测试更新权限不存在
		r3 := &biz.Permission{
			ID:   9999,
			Name: "test3",
			Memo: "test3",
		}
		_, err = permissionRepo.Update(context.Background(), r3)
		st, _ := status.FromError(err)
		convey.So(st.Code(), convey.ShouldEqual, codes.NotFound)

		// 测试更新权限名称重复
		r2.Path = "/test1"
		_, err = permissionRepo.Update(context.Background(), r2)
		st, _ = status.FromError(err)
		convey.So(st.Code(), convey.ShouldEqual, codes.AlreadyExists)
	}))
}

// TestDeletePermission 测试删除权限
func TestDeletePermission(t *testing.T) {
	convey.Convey("Test Delete Permission", t, run(func() {
		r := &biz.Permission{
			Name:   "test",
			Memo:   "test",
			Path:   "/test",
			Method: "GET",
		}
		newPermission, err := permissionRepo.Create(context.Background(), r)
		convey.So(err, convey.ShouldBeNil)
		convey.So(newPermission.Name, convey.ShouldEqual, r.Name)
		convey.So(newPermission.Memo, convey.ShouldEqual, r.Memo)
		convey.So(newPermission.Path, convey.ShouldEqual, r.Path)
		convey.So(newPermission.Method, convey.ShouldEqual, r.Method)
		convey.So(newPermission.ID, convey.ShouldNotEqual, 0)

		// 测试删除权限
		_, err = permissionRepo.Delete(context.Background(), newPermission.ID)
		convey.So(err, convey.ShouldBeNil)

		// 测试删除权限不存在
		_, err = permissionRepo.Delete(context.Background(), 9999)
		st, _ := status.FromError(err)
		convey.So(st.Code(), convey.ShouldEqual, codes.NotFound)
	}))
}

// TestGetPermission 测试获取权限
func TestGetPermission(t *testing.T) {
	convey.Convey("Test Get Permission", t, run(func() {
		r := &biz.Permission{
			Name:   "test",
			Memo:   "test",
			Path:   "/test",
			Method: "GET",
		}
		newPermission, err := permissionRepo.Create(context.Background(), r)
		convey.So(err, convey.ShouldBeNil)
		convey.So(newPermission.Name, convey.ShouldEqual, r.Name)
		convey.So(newPermission.Memo, convey.ShouldEqual, r.Memo)
		convey.So(newPermission.Path, convey.ShouldEqual, r.Path)
		convey.So(newPermission.Method, convey.ShouldEqual, r.Method)
		convey.So(newPermission.ID, convey.ShouldNotEqual, 0)

		// 测试获取权限
		where := map[string]interface{}{
			"id = ": newPermission.ID,
		}
		getPermission, err := permissionRepo.Get(context.Background(), where)
		convey.So(err, convey.ShouldBeNil)
		convey.So(getPermission.Name, convey.ShouldEqual, r.Name)
		convey.So(getPermission.Memo, convey.ShouldEqual, r.Memo)
		convey.So(getPermission.Path, convey.ShouldEqual, r.Path)
		convey.So(getPermission.Method, convey.ShouldEqual, r.Method)
		convey.So(getPermission.ID, convey.ShouldNotEqual, 0)

		// 测试获取权限不存在
		where = map[string]interface{}{
			"id = ": 9999,
		}
		_, err = permissionRepo.Get(context.Background(), where)
		st, _ := status.FromError(err)
		convey.So(st.Code(), convey.ShouldEqual, codes.NotFound)
	}))
}

// TestListPermission 测试获取权限列表
func TestListPermission(t *testing.T) {
	convey.Convey("Test List Permission", t, run(func() {
		for i := 0; i < 10; i++ {
			r := &biz.Permission{
				Name:   "test" + strconv.Itoa(i),
				Memo:   "test" + strconv.Itoa(i),
				Path:   "/test" + strconv.Itoa(i),
				Method: "GET",
			}
			newPermission, err := permissionRepo.Create(context.Background(), r)
			convey.So(err, convey.ShouldBeNil)
			convey.So(newPermission.Name, convey.ShouldEqual, r.Name)
			convey.So(newPermission.Memo, convey.ShouldEqual, r.Memo)
			convey.So(newPermission.Path, convey.ShouldEqual, r.Path)
			convey.So(newPermission.Method, convey.ShouldEqual, r.Method)
			convey.So(newPermission.ID, convey.ShouldNotEqual, 0)
		}

		// 测试获取权限列表
		where := map[string]interface{}{}
		order := map[string]bool{
			"id": true,
		}
		Permissions, total, err := permissionRepo.List(context.Background(), where, order, 0, 10)
		convey.So(err, convey.ShouldBeNil)
		convey.So(len(Permissions), convey.ShouldEqual, 1)
		convey.So(total, convey.ShouldEqual, 10)
	}))
}
