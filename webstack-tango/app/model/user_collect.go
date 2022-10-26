package model

import (
	"time"
)

type UserCollect struct {
	Id        uint64     `gorm:"primary_key;auto_increment;" json:"id"`
	UserId    uint64     `json:"user_id"`
	CollectId uint64     `json:"collect_id"`
	CreatedAt *time.Time `json:"created_at"` // 创建时间
}

func (UserCollect) TableName() string {
	return "user_collect"
}
