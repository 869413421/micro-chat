package data

import (
	"gorm.io/gorm"

	"github.com/869413421/micro-chat/pkg/db"
	"github.com/869413421/micro-chat/pkg/password"
)

// User 用户模型
type User struct {
	db.BaseModel
	Name     string `gorm:"column:name;type:varchar(255)  comment '用户名';not null;default:''" valid:"name"`
	Email    string `gorm:"column:email;type:varchar(50) comment '邮箱';not null;unique;default:''" valid:"email"`
	Password string `gorm:"column:password;type:varchar(255) comment '密码';not null;default:''" valid:"password"`
	RoleID   uint64 `gorm:"column:role_id;type:int comment '角色ID'; index;not null; default:0"`
}

// BeforeSave 保存前模型事件
func (model *User) BeforeSave(tx *gorm.DB) (err error) {
	//1.如果密码没加密，进行一次加密
	if !password.IsHashed(model.Password) {
		model.Password, err = password.Hash(model.Password)
		if err != nil {
			return err
		}
	}
	return nil
}
