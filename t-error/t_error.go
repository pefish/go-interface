package t_error

import (
	"fmt"

	"github.com/pkg/errors"
)

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

var (
	INTERNAL_ERROR_MSG         = "Internal error."
	INTERNAL_ERROR_CODE uint64 = 1
	NETWORK_ERROR_MSG          = "Network error."
	NETWORK_ERROR_CODE  uint64 = 2

	INTERNAL_ERROR = WrapWithCode(fmt.Errorf(INTERNAL_ERROR_MSG), INTERNAL_ERROR_CODE)
	NETWORK_ERROR  = WrapWithCode(fmt.Errorf(NETWORK_ERROR_MSG), NETWORK_ERROR_CODE)
)

func Recover(fun func(*ErrorInfo)) {
	if err := recover(); err != nil {
		switch e := err.(type) {
		case *ErrorInfo:
			fun(e)
		case error:
			fun(WrapWithAll(e, INTERNAL_ERROR_CODE, nil))
		default:
			fun(WrapWithAll(errors.New(fmt.Sprint(err)), INTERNAL_ERROR_CODE, nil))
		}
	}
}

func WithStack(err error) error {
	if _, ok := err.(interface {
		WithStack(err error) error
	}); !ok { // 如果是不带堆栈的err，则附加上堆栈
		return errors.WithStack(err)
	}
	return err
}

func ThrowInternal(err error) {
	var errorInfo_ = ErrorInfo{Code: INTERNAL_ERROR_CODE, Data: nil, Err: WithStack(err)}
	panic(&errorInfo_)
}

func ThrowWithCode(err error, code uint64) {
	var errorInfo_ = ErrorInfo{Code: code, Data: nil, Err: WithStack(err)}
	panic(&errorInfo_)
}

func ThrowWithCodeAndData(err error, code uint64, data error) {
	var errorInfo_ = ErrorInfo{Code: code, Data: data, Err: WithStack(err)}
	panic(&errorInfo_)
}

func WrapWithAll(err error, code uint64, data interface{}) *ErrorInfo {
	return &ErrorInfo{Code: code, Data: data, Err: WithStack(err)}
}

func WrapWithErr(err error) *ErrorInfo {
	return &ErrorInfo{Code: INTERNAL_ERROR_CODE, Data: nil, Err: WithStack(err)}
}

func WrapWithStr(str string) *ErrorInfo {
	return &ErrorInfo{Code: INTERNAL_ERROR_CODE, Data: nil, Err: WithStack(fmt.Errorf(str))}
}

func Wrap(err error) *ErrorInfo {
	return &ErrorInfo{Code: INTERNAL_ERROR_CODE, Data: nil, Err: WithStack(err)}
}

func WrapWithCode(err error, code uint64) *ErrorInfo {
	return &ErrorInfo{Code: code, Data: nil, Err: WithStack(err)}
}
