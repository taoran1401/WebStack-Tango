package model

import (
	"time"
)

type Collect struct {
	Id        uint64     `gorm:"primary_key;auto_increment;" json:"id"`
	Title     string     `json:"title"`      // 标题
	Desc      string     `json:"desc"`       // 描述
	Author    *int64     `json:"author"`     // 作者
	Link      string     `json:"link"`       // 链接
	Cover     string     `json:"cover"`      // 封面/logo
	CatId     uint64     `json:"cat_id"`     // 分类id
	Sort      int64      `json:"sort"`       // 权重，大的在前
	CreatedAt *time.Time `json:"created_at"` // 创建时间
	UpdatedAt *time.Time `json:"updated_at"` // 修改时间
	DeletedAt *time.Time `json:"deleted_at"` // 删除时间
}

func (Collect) TableName() string {
	return "collect"
}
