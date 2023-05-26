package schema

import (
	"github.com/869413421/micro-chat/app/user/service/internal/biz"
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

// ToBizUserRole 转换为biz层用户角色实体
func (u *UserRole) ToBizUserRole() *biz.UserRole {
	return &biz.UserRole{
		ID:     u.ID,
		UserID: u.UserID,
		RoleID: u.RoleID,
	}
}
