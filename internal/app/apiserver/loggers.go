package apiserver

import (
	"log"
	"os"
)

type Logger struct {
	InfoLog *log.Logger
	ErrLog  *log.Logger
}

func NewLogger() *Logger {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return &Logger{
		InfoLog: infoLog,
		ErrLog:  errorLog,
	}
}
