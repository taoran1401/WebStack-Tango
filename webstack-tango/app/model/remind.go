package model

import "time"

type Remind struct {
	Id        uint64     `gorm:"primary_key;auto_increment;" json:"id"`
	UserId    uint64     `json:"user_id"`
	Account   string     `json:"account"`
	Content   string     `json:"content"`
	Type      int64      `json:"type"`       //0 邮件； 1 公众号；
	CreatedAt *time.Time `json:"created_at"` // 创建时间
	UpdatedAt *time.Time `json:"updated_at"` // 修改时间
	DeletedAt *time.Time `json:"deleted_at"` // 删除时间
}

func (Remind) TableName() string {
	return "remind"
}
