package data

import "github.com/869413421/micro-chat/pkg/db"

type Permission struct {
	db.BaseModel
	Name      string `gorm:"column:name;type:varchar(255)  comment '权限名称';not null;default:''"`
	Memo      string `gorm:"column:memo;type:varchar(255)  comment '备注';not null;default:''"`
	Path      string `gorm:"column:path;type:varchar(255)  comment '路径';not null;default:''"`
	Method    string `gorm:"column:method;type:varchar(255)  comment '方法';not null;default:''"`
	ParentID  uint64 `gorm:"column:parent_id;type:int comment '父级ID'; index;not null; default:0"`
	ParentIDS string `gorm:"column:parent_ids;type:varchar(255)  comment '父级路径';not null;default:''"`
}
