package routers

import "maya/pkg/errcode"

type ResResult[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func (r ResResult[T]) GetErrorResResult(code int) ResResult[T] {
	r.Code = code
	r.Message = errcode.GetDesc(code)
	return r
}

type ResPageResult[T any] struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	TotalCount int    `json:"totalCount"`
	Data       []T    `json:"data"`
}

func NewResResult[T any]() ResResult[T] {
	return ResResult[T]{
		Code: errcode.Failed,
	}
}
