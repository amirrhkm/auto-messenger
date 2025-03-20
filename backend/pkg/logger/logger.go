package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

type Logger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	SetBroadcastFunc(fn func(string))
}

type logger struct {
	log           *log.Logger
	ws            *websocket.Conn
	broadcastFunc func(string)
}

func NewLogger(level string) Logger {
	return &logger{
		log: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	}
}

func (l *logger) SetBroadcastFunc(fn func(string)) {
	l.broadcastFunc = fn
}

func (l *logger) Info(msg string, args ...interface{}) {
	formattedMsg := fmt.Sprintf("[INFO] "+msg, args...)
	l.log.Print(formattedMsg)

	if l.broadcastFunc != nil {
		l.broadcastFunc(formattedMsg)
	}
}

func (l *logger) Error(msg string, args ...interface{}) {
	formattedMsg := fmt.Sprintf("[ERROR] "+msg, args...)
	l.log.Print(formattedMsg)

	if l.broadcastFunc != nil {
		l.broadcastFunc(formattedMsg)
	}
}
