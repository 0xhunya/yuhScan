package logger

import (
	"log"
	"os"
	"syscall"
)

type logType struct {
	color  int
	prefix string
}

const (
	BLACK = iota
	BLUE
	GREEN
	CYAN
	RED
	PURPLE
	YELLOW
	LIGHT_GRAY
	GRAY
	LIGHT_BLUE
	LIGHT_GREEN
	LIGHT_CYAN
	LIGHT_RED
	LIGHT_PURPLE
	LIGHT_YELLOW
	WHITE
)

var (
	STATUS = logType{color: GRAY, prefix: "[STATUS]"}
	INFO   = logType{color: LIGHT_GREEN, prefix: "[INFO]"}
	WARN   = logType{color: LIGHT_YELLOW, prefix: "[WARN]"}
	ERROR  = logType{color: LIGHT_RED, prefix: "[ERROR]"}
	DEBUG  = logType{color: YELLOW, prefix: "[DEBUG]"}
	API    = logType{color: LIGHT_GRAY, prefix: "[API]"}
	NORMAL = logType{color: GRAY, prefix: "[+]"}
	RESULT = logType{color: WHITE, prefix: "=============== RESULT ===============\n"}
)

func ConsoleLog(logt logType, v ...interface{}) {
	var (
		loggerTime *log.Logger = log.New(os.Stdout, logt.prefix+" ", log.Ldate|log.Ltime)
		logger     *log.Logger = log.New(os.Stdout, logt.prefix, 0)

		kernel32    *syscall.LazyDLL  = syscall.NewLazyDLL(`kernel32.dll`)
		proc        *syscall.LazyProc = kernel32.NewProc(`SetConsoleTextAttribute`)
		CloseHandle *syscall.LazyProc = kernel32.NewProc(`CloseHandle`)
	)

	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(logt.color))

	switch logt {
	case RESULT:
		logger.Println(v...)
	case ERROR:
		logger.Fatalln(v...)
	default:
		loggerTime.Println(v...)
	}

	CloseHandle.Call(handle)
}
