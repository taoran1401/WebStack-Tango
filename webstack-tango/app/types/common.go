package types

type PageResult struct {
	List     interface{} `json:"list"`
	Total    uint64      `json:"total"`
	Page     uint64      `json:"page"`
	PageSize uint64      `json:"page_size"`
}
