package t_error

type ErrorInfo struct {
	Code uint64      `json:"code"`
	Data interface{} `json:"data"`
	Err  error       `json:"err"`
}
