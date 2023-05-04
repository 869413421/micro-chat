package data_test

import (
	"context"
	"github.com/869413421/micro-chat/app/user/service/internal/biz"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TestCreateRole 测试创建角色
func TestCreateRole(t *testing.T) {
	convey.Convey("Test Create Role", t, run(func() {
		r := &biz.Role{
			Name: "test",
			Memo: "test",
		}
		newRole, err := roleRepo.Create(context.Background(), r)
		convey.So(err, convey.ShouldBeNil)
		convey.So(newRole.Name, convey.ShouldEqual, r.Name)
		convey.So(newRole.ID, convey.ShouldNotEqual, 0)

		// Test error case when server already exists
		_, err = roleRepo.Create(context.Background(), r)
		st, _ := status.FromError(err)

		convey.So(st.Code(), convey.ShouldEqual, codes.AlreadyExists)
	}))
}

// TestUpdateRole 测试更新角色
func TestUpdateRole(t *testing.T) {
	convey.Convey("Test Update Role", t, run(func() {
		r := &biz.Role{
			Name: "test",
			Memo: "test",
		}
		newRole, err := roleRepo.Create(context.Background(), r)
		convey.So(err, convey.ShouldBeNil)
		convey.So(newRole.Name, convey.ShouldEqual, r.Name)
		convey.So(newRole.ID, convey.ShouldNotEqual, 0)

		// 测试更新角色名称
		r.ID = newRole.ID
		r.Name = "test2"
		newRole, err = roleRepo.Update(context.Background(), r)
		convey.So(err, convey.ShouldBeNil)
		convey.So(newRole.Name, convey.ShouldEqual, "test2")
	}))
}

// TestUpdateRole_Err 测试更新角色错误
func TestUpdateRole_Err(t *testing.T) {
	convey.Convey("Test Update Role", t, run(func() {
		r := &biz.Role{
			Name: "test1",
			Memo: "test1",
		}
		newRole, err := roleRepo.Create(context.Background(), r)
		convey.So(err, convey.ShouldBeNil)
		convey.So(newRole.Name, convey.ShouldEqual, r.Name)
		convey.So(newRole.ID, convey.ShouldNotEqual, 0)

		r2 := &biz.Role{
			Name: "test2",
			Memo: "test2",
		}
		newRole2, err := roleRepo.Create(context.Background(), r2)
		convey.So(err, convey.ShouldBeNil)
		convey.So(newRole2.Name, convey.ShouldEqual, r2.Name)
		convey.So(newRole2.ID, convey.ShouldNotEqual, 0)

		// 测试更新角色不存在
		r3 := &biz.Role{
			ID:   9999,
			Name: "test3",
			Memo: "test3",
		}
		_, err = roleRepo.Update(context.Background(), r3)
		st, _ := status.FromError(err)
		convey.So(st.Code(), convey.ShouldEqual, codes.NotFound)

		// 测试更新角色名称重复
		r2.Name = "test1"
		_, err = roleRepo.Update(context.Background(), r2)
		st, _ = status.FromError(err)
		convey.So(st.Code(), convey.ShouldEqual, codes.AlreadyExists)
	}))
}

// TestDeleteRole 测试删除角色
func TestDeleteRole(t *testing.T) {
	convey.Convey("Test Delete Role", t, run(func() {
		r := &biz.Role{
			Name: "test",
			Memo: "test",
		}
		newRole, err := roleRepo.Create(context.Background(), r)
		convey.So(err, convey.ShouldBeNil)
		convey.So(newRole.Name, convey.ShouldEqual, r.Name)
		convey.So(newRole.ID, convey.ShouldNotEqual, 0)

		// 测试删除角色
		_, err = roleRepo.Delete(context.Background(), newRole.ID)
		convey.So(err, convey.ShouldBeNil)

		// 测试删除角色不存在
		_, err = roleRepo.Delete(context.Background(), 9999)
		st, _ := status.FromError(err)
		convey.So(st.Code(), convey.ShouldEqual, codes.NotFound)
	}))
}

// TestGetRole 测试获取角色
func TestGetRole(t *testing.T) {
	convey.Convey("Test Get Role", t, run(func() {
		r := &biz.Role{
			Name: "test",
			Memo: "test",
		}
		newRole, err := roleRepo.Create(context.Background(), r)
		convey.So(err, convey.ShouldBeNil)
		convey.So(newRole.Name, convey.ShouldEqual, r.Name)
		convey.So(newRole.ID, convey.ShouldNotEqual, 0)

		// 测试获取角色
		where := map[string]interface{}{
			"id = ": newRole.ID,
		}
		role, err := roleRepo.Get(context.Background(), where)
		convey.So(err, convey.ShouldBeNil)
		convey.So(role.Name, convey.ShouldEqual, r.Name)
		convey.So(role.Memo, convey.ShouldEqual, r.Memo)
		convey.So(role.ID, convey.ShouldEqual, newRole.ID)

		// 测试获取角色不存在
		where = map[string]interface{}{
			"id = ": 9999,
		}
		_, err = roleRepo.Get(context.Background(), where)
		st, _ := status.FromError(err)
		convey.So(st.Code(), convey.ShouldEqual, codes.NotFound)
	}))
}

// TestListRole 测试获取角色列表
func TestListRole(t *testing.T) {
	convey.Convey("Test List Role", t, run(func() {
		r := &biz.Role{
			Name: "test",
			Memo: "test",
		}
		newRole, err := roleRepo.Create(context.Background(), r)
		convey.So(err, convey.ShouldBeNil)
		convey.So(newRole.Name, convey.ShouldEqual, r.Name)
		convey.So(newRole.ID, convey.ShouldNotEqual, 0)

		// 测试获取角色列表
		where := map[string]interface{}{
			"id = ": newRole.ID,
		}
		order := map[string]bool{
			"id": true,
		}
		roles, total, err := roleRepo.List(context.Background(), where, order, 0, 10)
		convey.So(err, convey.ShouldBeNil)
		convey.So(len(roles), convey.ShouldEqual, 1)
		convey.So(total, convey.ShouldEqual, 1)
		convey.So(roles[0].Name, convey.ShouldEqual, r.Name)
		convey.So(roles[0].Memo, convey.ShouldEqual, r.Memo)
		convey.So(roles[0].ID, convey.ShouldEqual, newRole.ID)
	}))
}
