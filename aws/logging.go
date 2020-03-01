package aws

import (
	"log"
	"os"
)

type logLevel int

const (
	debug logLevel = iota
	info
	err
	fatal
)

var globalLogLevel = fatal

func setGlobalLogLevel() {
	l := os.Getenv("LOG_LEVEL")
	switch l {
	case "DEBUG":
		globalLogLevel = debug
	case "INFO":
		globalLogLevel = info
	case "ERROR":
		globalLogLevel = err
	case "FATAL":
		globalLogLevel = fatal
	default:
		globalLogLevel = fatal
	}
}

func logDebug(s ...interface{}) {
	if debug < globalLogLevel {
		return
	}
	log.Println("[DEBUG]", s)
}

func logInfo(s ...interface{}) {
	if info < globalLogLevel {
		return
	}
	log.Println("[INFO]", s)
}

func logError(s ...interface{}) {
	if err < globalLogLevel {
		return
	}
	log.Println("[ERROR]", s)
}

func logFatal(s ...interface{}) {
	if fatal < globalLogLevel {
		return
	}
	log.Fatalln("[FATAL]", s)
}
