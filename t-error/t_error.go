package t_error

import "fmt"

type ErrorInfo struct {
	Code uint64      `json:"code"`
	Data interface{} `json:"data"`
	Err  error       `json:"err"`
}

func (ei *ErrorInfo) String() string {
	return fmt.Sprintf(`ErrorInfo -> msg: %s, code: %d, data: %#v, err: %+v`, ei.Err.Error(), ei.Code, ei.Data, ei.Err)
}

func (ei *ErrorInfo) Error() string {
	return ei.String()
}

func (ei *ErrorInfo) Unwrap() error {
	return ei.Err
}
