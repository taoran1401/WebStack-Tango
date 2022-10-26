package model

import (
	"time"
)

type CollectCategory struct {
	Id        uint64     `gorm:"primary_key;auto_increment;" json:"id"`
	Name      string     `json:"name"`       // 名称
	PId       int64      `json:"p_id"`       // 上级id
	Sort      int64      `json:"sort"`       // 权重，大的在前
	CreatedAt *time.Time `json:"created_at"` // 创建时间
	UpdatedAt *time.Time `json:"updated_at"` // 修改时间
	DeletedAt *time.Time `json:"deleted_at"` // 删除时间
	Collect   []Collect  `gorm:"foreignkey:CatId;" json:"collect"`
}

func (CollectCategory) TableName() string {
	return "collect_category"
}
