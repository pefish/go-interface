package t_mysql

import "time"

type Configuration struct {
	Host            string
	Port            int
	Username        string
	Password        string
	Database        string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnParams      map[string]string
}

type CountParams struct {
	TableName string
	Where     interface{}
}

type SumParams struct {
	TableName string
	SumTarget string
	Where     interface{}
}

type OrderType string

const (
	OrderType_ASC  OrderType = "asc"
	OrderType_DESC OrderType = "desc"
)

type OrderByType struct {
	Col   string
	Order OrderType
}

type SelectParams struct {
	TableName string
	Select    string
	Where     interface{}
	OrderBy   *OrderByType
	Limit     uint64
}

type SelectByIdParams struct {
	TableName string
	Select    string
	Id        uint64
}

type UpdateParams struct {
	TableName string
	Update    interface{}
	Where     interface{}
}
