package schema

import (
	"github.com/869413421/micro-chat/app/user/service/internal/biz"
	"gorm.io/gorm"

	"github.com/869413421/micro-chat/pkg/db"
	"github.com/869413421/micro-chat/pkg/password"
)

// User 用户模型
type User struct {
	db.BaseModel
	Name      string      `gorm:"column:name;type:varchar(255)  comment '用户名';not null;default:''"`
	Email     string      `gorm:"index:,unique uniqueIndex unique;column:email;type:varchar(50) comment '邮箱';not null;default:''"`
	Password  string      `gorm:"column:password;type:varchar(255) comment '密码';not null;default:''" `
	Status    int         `gorm:"column:status;type:tinyint comment '状态';not null;default:1"`
	UserRoles []*UserRole `gorm:"foreignKey:UserID;references:ID"`
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

// ToBizUser 转换为biz层用户实体
func (model *User) ToBizUser() *biz.User {
	return &biz.User{
		ID:       model.ID,
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
		Status:   model.Status,
		CreateAt: model.CreatedAtDate(),
		UpdateAt: model.UpdatedAtDate(),
	}
}
