package data

import "github.com/869413421/micro-chat/pkg/db"

// Role 角色模型
type Role struct {
	db.BaseModel
	Name            string            `gorm:"column:name;type:varchar(255)  comment '角色名称';not null;unique;default:''"`
	Memo            string            `gorm:"column:memo;type:varchar(255)  comment '备注';not null;default:''"`
	RolePermissions []*RolePermission `gorm:"foreignKey:RoleID;references:ID"`
}
