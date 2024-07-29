package i_mysql

import (
	t_mysql "github.com/pefish/go-interface/t-mysql"
)

type IMysql interface {
	TagName() string
	Close()

	ConnectWithConfiguration(configuration t_mysql.Configuration) error

	RawSelectFirst(
		dest interface{},
		select_ string,
		str string,
		values ...interface{},
	) (
		notFound bool,
		err error,
	)

	RawSelect(
		dest interface{},
		select_ string,
		str string,
		values ...interface{},
	) error

	RawExec(sql string, values ...interface{}) (
		lastInsertId uint64,
		err error,
	)
	Count(countParams *t_mysql.CountParams, values ...interface{}) (
		count uint64,
		err error,
	)
	RawCount(sql string, values ...interface{}) (
		count uint64,
		err error,
	)
	Sum(sumParams *t_mysql.SumParams, values ...interface{}) (
		sum float64,
		err error,
	)
	SelectFirst(dest interface{}, selectParams *t_mysql.SelectParams, values ...interface{}) (
		notFound bool,
		err error,
	)
	SelectById(
		dest interface{},
		selectByIdParams *t_mysql.SelectByIdParams,
	) (
		notFound bool,
		err error,
	)
	Select(dest interface{}, selectParams *t_mysql.SelectParams, values ...interface{}) error
	Insert(tableName string, params interface{}) (
		lastInsertId uint64,
		err error,
	)
	Update(
		updateParams *t_mysql.UpdateParams,
		values ...interface{},
	) (
		lastInsertId uint64,
		err error,
	)

	Begin() (IMysql, error)
	Commit() error
	Rollback() error
}
