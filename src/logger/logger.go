package logger

import (
	"log"
	"os"
	"io"
)

var (
	info           *log.Logger
	warning        *log.Logger
	error          *log.Logger
)

func init() {
	logFile, err := os.OpenFile(os.Getenv("LOG_PATH"), os.O_CREATE | os.O_APPEND | os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	info = log.New(io.MultiWriter(os.Stdout, logFile), "INFO: ", log.Ldate|log.Ltime)
	warning = log.New(io.MultiWriter(os.Stdout, logFile), "WARNING: ", log.Ldate|log.Ltime)
	error = log.New(io.MultiWriter(os.Stdout, logFile), "ERROR: ", log.Ldate|log.Ltime)
}

func Info(v ...interface{})  {
	info.Println(v)
}

func Warning(v ...interface{}) {
	warning.Println(v)
}

func Error(v ...interface{})  {
	error.Println(v)
}