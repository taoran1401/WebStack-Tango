package types

import "time"

type CollectBase struct {
	Id    uint64 `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	CatId uint64 `json:"cat_id"`
	Cover string `json:"cover"`
	Link  string `json:"link"`
	Sort  int64  `json:"sort"`
}

type CollectListReq struct {
	Title    string `json:"title"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"page_size"`
}

type CollectListResp struct {
	Page     uint64
	PageSize uint64
	Total    uint64
	List     []map[string]string
}

type SaveCollectReq struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
	CatId uint64 `json:"cat_id"`
	Cover string `json:"cover"`
	Link  string `json:"link"`
	Sort  int64  `json:"sort"`
}

type CollectStarResp struct {
	Id        uint64     `json:"id"`
	UserId    uint64     `json:"user_id"`
	CollectId uint64     `json:"collect_id"`
	CreatedAt *time.Time `json:"created_at"` // 创建时间
}

type CollectGetStarReq struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"page_size"`
}

type CollectGetStarResp struct {
	Page     uint64
	PageSize uint64
	Total    uint64
	List     []map[string]string
}
