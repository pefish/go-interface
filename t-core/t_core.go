package t_core

type ApiResult struct {
	Msg  string      `json:"msg"`
	Code uint64      `json:"code"`
	Data interface{} `json:"data"`
}

type StatusCode int
