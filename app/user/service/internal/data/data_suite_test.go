package data_test

import (
	"github.com/869413421/micro-chat/app/user/service/internal/data/orm/schema"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

	"github.com/869413421/micro-chat/app/user/service/internal/biz"
	"github.com/869413421/micro-chat/app/user/service/internal/conf"
	"github.com/869413421/micro-chat/app/user/service/internal/data"
)

var userRepo biz.UserRepo
var roleRepo biz.RoleRepo
var permissionRepo biz.PermissionRepo
var db *gorm.DB

// init 初始化测试依赖项
func init() {
	config := &conf.Data{Database: &conf.Data_Database{Driver: "mysql", Source: "root:root@tcp(user-db:33061)/user?charset=utf8mb4&parseTime=True&loc=Local"}}
	var err error
	db, err = data.NewDB(config)
	if err != nil {
		log.Fatal(err)
	}

	newData, _, _ := data.NewData(config, nil, db)
	userRepo, _ = data.NewUserRepo(newData, nil)
	roleRepo, _ = data.NewRoleRepo(newData, nil)
	permissionRepo, _ = data.NewPermissionRepo(newData, nil)
}

// run 闭包用于删除测试数据
func run(f func()) func() {
	return func() {
		defer func() {
			db.Unscoped().Where("id > ?", 0).Delete(&schema.User{})
			db.Unscoped().Where("id > ?", 0).Delete(&schema.Role{})
			db.Unscoped().Where("id > ?", 0).Delete(&schema.Permission{})
		}()
		f()
	}
}
