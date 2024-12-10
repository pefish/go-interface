package i_core

import (
	"io"
	"net/http"

	i_logger "github.com/pefish/go-interface/i-logger"
	t_core "github.com/pefish/go-interface/t-core"
	t_error "github.com/pefish/go-interface/t-error"
)

type IApi interface {
	Description() string
	ParamType() string
	Params() interface{}
}

type IApiSession interface {
	SetPathVars(vars map[string]string)
	PathVars() map[string]string
	SetJwtBody(jwtBody map[string]interface{})
	JwtBody() map[string]interface{}
	SetUserId(userId uint64)
	UserId() uint64
	SetJwtHeaderName(headerName string)
	JwtHeaderName() string
	ScanParams(dest interface{}) error
	MustScanParams(dest interface{})
	AddDefer(defer_ func())
	Defers() []func()
	SetData(key string, data interface{})
	Data(key string) interface{}
	Redirect(url string)
	WriteJson(data interface{}) error
	SetHeader(key string, value string)
	WriteText(text string) error
	SetStatusCode(code t_core.StatusCode)
	Host() string
	Path() string
	Body() io.ReadCloser
	Method() string
	Header(name string) string
	RemoteAddress() string
	UrlParams() map[string]string
	FormValues() (map[string][]string, error)
	ReadJSON(jsonObject interface{}) error
	ReadMap() (map[string]interface{}, error)
	Api() IApi
	SetApi(api IApi)
	ResponseWriter() http.ResponseWriter
	SetResponseWriter(w http.ResponseWriter)
	Request() *http.Request
	SetRequest(r *http.Request)
	Params() map[string]interface{}
	SetParams(params map[string]interface{})
	OriginalParams() map[string]interface{}
	SetOriginalParams(originalParams map[string]interface{})
	SetLang(lang string)
	Lang() string
	SetClientType(clientType string)
	ClientType() string
	Logger() i_logger.ILogger
}

type IApiStrategy interface {
	Execute(out IApiSession) *t_error.ErrorInfo
	Name() string
	Description() string
	ErrorCode() uint64
	SetErrorCode(code uint64) IApiStrategy
	SetErrorMsg(msg string) IApiStrategy
	ErrorMsg() string
}
