package types

type RemindBase struct {
	Id      uint64 `json:"id"`
	UserId  uint64 `json:"user_id"`
	Account string `json:"account"`
	Content string `json:"content"`
	Type    int64  `json:"type"` //0 邮件； 1 公众号；
}

type SaveRemindReq struct {
	Account string `json:"account"`
	Content string `json:"content"`
	Type    int64  `json:"type"` //0 邮件； 1 公众号；
}
