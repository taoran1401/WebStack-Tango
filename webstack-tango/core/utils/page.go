package utils

import "github.com/jinzhu/gorm"

type PageInfo struct {
	Total    uint64      `json:"total"`
	Page     uint64      `json:"page"`
	PageSize uint64      `json:"page_size"`
	List     interface{} `json:"list"`
}

//分页器
//db: gorm db
//page: 页码
//pageSize: 每页数量
//list: list
func Paginate(db *gorm.DB, page uint64, pageSize uint64, list interface{}) (PageInfo, error) {
	//获取总数
	var total uint64
	err := db.Count(&total).Error
	if err != nil {
		return PageInfo{}, err
	}

	//获取列表
	offset := (page - 1) * pageSize
	db.Offset(offset).Limit(pageSize).Find(&list)
	return PageInfo{
		Total:    total,
		Page:     page,
		PageSize: pageSize,
		List:     list,
	}, nil
}
