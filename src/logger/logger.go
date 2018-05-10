package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	info    *log.Logger
	warning *log.Logger
	erro    *log.Logger
	debug   *log.Logger
)

func init() {
	logFile, err := os.OpenFile(os.Getenv("LOG_PATH"), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)

	if err != nil {
		panic(fmt.Sprintf("Cant open/create log file by path: %v", os.Getenv("LOG_PATH")))
	}

	info = log.New(io.MultiWriter(os.Stdout, logFile), "INFO: ", log.Ldate|log.Ltime)
	warning = log.New(io.MultiWriter(os.Stdout, logFile), "WARNING: ", log.Ldate|log.Ltime)
	erro = log.New(io.MultiWriter(os.Stdout, logFile), "ERROR: ", log.Ldate|log.Ltime)
	debug = log.New(io.MultiWriter(os.Stdout, logFile), "DEBUG: ", log.Ldate|log.Ltime)
}

func Debug(v ...interface{}) {
	if os.Getenv("DEBUG") == "true" {
		debug.Println(v)
	}
}

func Info(v ...interface{}) {
	info.Println(v)
}

func Warning(v ...interface{}) {
	warning.Println(v)
}

func Error(v ...interface{}) {
	erro.Fatal(v)
}
