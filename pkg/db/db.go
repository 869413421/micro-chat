package db

import (
	"gorm.io/gorm"
	"strconv"
	"time"
)

type BaseModel struct {
	ID        uint64         "gorm:column:id;primaryKey;autoIncrement;not null"
	CreatedAt time.Time      `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"column:updated_at;index;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

//GetStringID 主键转字符串
func (model BaseModel) GetStringID() string {
	return strconv.Itoa(int(model.ID))
}

// CreatedAtDate 获取模型创建时间
func (model BaseModel) CreatedAtDate() string {
	return model.CreatedAt.Format("2006-01-02 15:04:05")
}

// UpdatedAtDate 获取模型更新时间
func (model BaseModel) UpdatedAtDate() string {
	return model.UpdatedAt.Format("2006-01-02 15:04:05")
}
