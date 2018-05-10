package logger

import (
	"log"
	"io/ioutil"
	"os"
)

var (
	trace          *log.Logger
	info           *log.Logger
	warning        *log.Logger
	error          *log.Logger
)

func init() {
	trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Trance(v ...interface{})  {
	trace.Println(v)
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