package i_logger

import t_logger "github.com/pefish/go-interface/t-logger"

type ILogger interface {
	FormatOutput(args ...interface{}) string
	Sdump(args ...interface{}) string
	Level() t_logger.Level
	CloneWithPrefix(prefix string) *ILogger
	CloneWithLevel(level t_logger.Level) *ILogger
	CloneWithOutputFile(filepath string) *ILogger

	Debug(args ...interface{})
	DebugF(format string, args ...interface{})
	DebugFRaw(format string, args ...interface{})

	Info(args ...interface{})
	InfoF(format string, args ...interface{})
	InfoFRaw(format string, args ...interface{})
	InfoDump(args ...interface{})

	Warn(args ...interface{})
	WarnF(format string, args ...interface{})
	WarnFRaw(format string, args ...interface{})

	Error(args ...interface{})
	ErrorF(format string, args ...interface{})
	ErrorFRaw(format string, args ...interface{})
}
