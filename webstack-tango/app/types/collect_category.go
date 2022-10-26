package types

type CollectCategoryBase struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	PId  int64  `json:"p_id"`
	Sort int64  `json:"sort"`
}

type CollectCategoryListReq struct {
	Name     string `json:"name"`
	Page     int64  `json:"page"`
	PageSize int64  `json:"page_size"`
}

type CollectCategoryListResp struct {
	Page     uint64
	PageSize uint64
	Total    uint64
	List     []map[string]string
}

type SaveCollectCategoryReq struct {
	Name string `json:"name"`
	Sort int64  `json:"sort"`
}
