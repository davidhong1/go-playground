package log

import (
	"fmt"
	"runtime"
)

func INFO(format string, args ...interface{}) {
	log("INFO", format, args...)
}

func ERROR(format string, args ...interface{}) {
	log("ERROR", format, args...)
}

func log(level string, format string, args ...interface{}) {
	filename, line, funcName := "???", 0, "???"
	pc, filename, line, ok := runtime.Caller(2)
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}

	fmt.Printf("%5s: %s:%d:%s: %s\n", level, filename, line, funcName, fmt.Sprintf(format, args...))
}
