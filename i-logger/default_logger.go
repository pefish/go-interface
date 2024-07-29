package i_logger

import "fmt"

var DefaultLogger = DefaultLoggerType{}

type DefaultLoggerType struct{}

func (d *DefaultLoggerType) FormatOutput(args ...interface{}) string {
	if len(args) == 0 {
		return ""
	}
	result := ``
	for _, arg := range args {
		result += fmt.Sprintf("%+v", arg) + "   "
	}
	result = result[:len(result)-3]
	return result
}

func (d *DefaultLoggerType) Sdump(args ...interface{}) string {
	return d.FormatOutput(args...)
}

func (d *DefaultLoggerType) Debug(args ...interface{}) {
	fmt.Println(args...)
}

func (d *DefaultLoggerType) DebugF(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (d *DefaultLoggerType) DebugFRaw(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (d *DefaultLoggerType) Info(args ...interface{}) {
	fmt.Println(args...)
}

func (d *DefaultLoggerType) InfoF(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (d *DefaultLoggerType) InfoFRaw(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (d *DefaultLoggerType) InfoDump(args ...interface{}) {
	fmt.Println(args...)
}

func (d *DefaultLoggerType) Warn(args ...interface{}) {
	fmt.Println(args...)
}

func (d *DefaultLoggerType) WarnF(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (d *DefaultLoggerType) WarnFRaw(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (d *DefaultLoggerType) Error(args ...interface{}) {
	fmt.Println(args...)
}

func (d *DefaultLoggerType) ErrorF(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (d *DefaultLoggerType) ErrorFRaw(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
