package types

type CollectCategoryBase struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	PId  int64  `json:"p_id"`
	Sort int64  `json:"sort"`
}

type CollectCategoryListReq struct {
	Name     string `json:"name"`
	Page     uint64 `json:"page"`
	PageSize uint64 `json:"page_size"`
}

type SaveCollectCategoryReq struct {
	Name string `json:"name"`
	Sort int64  `json:"sort"`
}
