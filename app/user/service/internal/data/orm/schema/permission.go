package schema

import (
	"github.com/869413421/micro-chat/app/user/service/internal/biz"
	"github.com/869413421/micro-chat/pkg/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"strings"

	"github.com/869413421/micro-chat/pkg/db"
)

type Permission struct {
	db.BaseModel
	Name      string `gorm:"column:name;type:varchar(255)  comment '权限名称';not null;default:''"`
	Memo      string `gorm:"column:memo;type:varchar(255)  comment '备注';not null;default:''"`
	Path      string `gorm:"column:path;type:varchar(255)  comment '路径';not null;default:''"`
	Method    string `gorm:"column:method;type:varchar(255)  comment '方法';not null;default:''"`
	ParentID  uint64 `gorm:"column:parent_id;type:int comment '父级ID'; index:parent_id;not null; default:0"`
	ParentIDS string `gorm:"column:parent_ids;type:varchar(255)  comment '父级路径';not null;default:''"`
}

// BeforeSave 保存前模型事件
func (model *Permission) BeforeSave(tx *gorm.DB) (err error) {
	if model.ParentID == 0 {
		return nil
	}

	var parent Permission
	result := tx.Where("id = ?", model.ParentID).First(&parent)
	if result.RowsAffected == 0 {
		return status.Errorf(codes.NotFound, "父级权限不存在")
	}
	model.ParentIDS = strings.Trim(parent.ParentIDS+","+types.UInt64ToString(model.ParentID), ",")
	return nil
}

// ToBizPermission 转换为biz层权限实体
func (model *Permission) ToBizPermission() *biz.Permission {
	return &biz.Permission{
		ID:        model.ID,
		Name:      model.Name,
		Memo:      model.Memo,
		Path:      model.Path,
		Method:    model.Method,
		ParentID:  model.ParentID,
		ParentIDS: model.ParentIDS,
		CreateAt:  model.CreatedAtDate(),
		UpdateAt:  model.UpdatedAtDate(),
	}
}
