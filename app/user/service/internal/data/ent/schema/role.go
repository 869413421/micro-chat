package schema

import (
	"github.com/869413421/micro-chat/app/user/service/internal/biz"
	"github.com/869413421/micro-chat/pkg/db"
)

// Role 角色模型
type Role struct {
	db.BaseModel
	Name            string            `gorm:"index:,unique uniqueIndex unique;column:name;type:varchar(255)  comment '角色名称';not null;default:''"` //定义唯一索引要按当前格式，否则会重复创建，参考issue：https://github.com/go-gorm/gorm/issues/5752。
	Memo            string            `gorm:"column:memo;type:varchar(255)  comment '备注';not null;default:''"`
	RolePermissions []*RolePermission `gorm:"foreignKey:RoleID;references:ID"`
}

// ToBizRole 转换为biz层角色实体
func (model *Role) ToBizRole() *biz.Role {
	return &biz.Role{
		ID:       model.ID,
		Name:     model.Name,
		Memo:     model.Memo,
		CreateAt: model.CreatedAtDate(),
		UpdateAt: model.UpdatedAtDate(),
	}
}
