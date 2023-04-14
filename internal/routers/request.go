package routers

type ReqPageCond[T any] struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
	Cond     T   `json:"cond"`
}
