package types

import (
	"taogin/app/model"
)

type IndexResp struct {
	Page     uint64
	PageSize uint64
	Total    uint64
	List     []IndexBase
}

type IndexBase struct {
	CollectCategory []model.CollectCategory
	/*Id      uint64 `gorm:"primary_key;auto_increment;" json:"id"`
	Name    string `json:"name"` // 名称
	PId     string `json:"p_id"` // 上级id
	Sort    int64  `json:"sort"` // 权重，大的在前*/
	//Collect []model.Collect
}
