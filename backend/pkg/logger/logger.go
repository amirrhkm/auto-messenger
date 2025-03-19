package logger

import (
	"log"
	"os"
)

type Logger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

type logger struct {
	log *log.Logger
}

func NewLogger(level string) Logger {
	return &logger{
		log: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	}
}

func (l *logger) Info(msg string, args ...interface{}) {
	l.log.Printf("[INFO] "+msg, args...)
}

func (l *logger) Error(msg string, args ...interface{}) {
	l.log.Printf("[ERROR] "+msg, args...)
}
