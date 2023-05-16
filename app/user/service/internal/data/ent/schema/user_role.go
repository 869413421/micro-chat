package schema

import (
	"github.com/869413421/micro-chat/pkg/db"
)

// UserRole 用户角色
type UserRole struct {
	db.BaseModel
	UserID uint64 `gorm:"column:user_id;type:int comment '用户ID'; index;not null; default:0"`
	User   *User
	RoleID uint64 `gorm:"column:role_id;type:int comment '角色ID'; index;not null; default:0"`
	Role   *Role
}
