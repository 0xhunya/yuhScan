package logger

import (
	"log"
	"os"
)

const (
	// Color
	WHITE            = "\033[37m"
	GREEN            = "\033[92m"
	RED              = "\033[91m"
	YELLOW           = "\033[33m"
	HIGHLIGHT_YELLOW = "\033[93m"
	BLUE             = "\033[94m"
	HIGHLIGHT_WHITE  = "\033[97m"
	ENDC             = "\033[0m"
	// Type
	STATUS = WHITE + "[STATUS]"
	INFO   = GREEN + "[INFO]"
	WARN   = HIGHLIGHT_YELLOW + "[WARN]"
	ERROR  = RED + "[ERROR]"
	DEBUG  = YELLOW + "[DEBUG]"
	API    = HIGHLIGHT_WHITE + "[API]"
	NORMAL = WHITE + "[+]"
	RESULT = HIGHLIGHT_WHITE + "=============== RESULT ===============" + "\n"
)

func ConsoleLog(prefix string, v ...interface{}) {
	loggerTime := log.New(os.Stdout, prefix+" ", log.Ldate|log.Ltime)
	logger := log.New(os.Stdout, prefix, 0)
	switch prefix {
	case RESULT:
		logger.Println(v...)
	case ERROR:
		logger.Fatalln(v...)
	default:
		loggerTime.Println(v...)
	}
	print(ENDC)
}
