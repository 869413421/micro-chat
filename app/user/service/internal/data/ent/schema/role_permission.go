package schema

import (
	"github.com/869413421/micro-chat/pkg/db"
)

// RolePermission 角色权限
type RolePermission struct {
	db.BaseModel        // 唯一标识
	RoleID       uint64 `gorm:"column:role_id;type:int comment '角色ID'; index;not null; default:0"`
	Role         *Role
	PermissionID uint64 `gorm:"column:permission_id;type:int comment '权限ID'; index;not null; default:0"`
	Permission   *Permission
}
