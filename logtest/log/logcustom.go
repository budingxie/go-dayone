package log11

import (
	"fmt"
	"log"
)

const (
	LevelError = iota
	LevelWarning
	LevelInformational
	LevelDebug
)

type Logger struct {
	level int
	l     *log.Logger
}

func (ll *Logger) Error(format string, v ...interface{}) {
	if LevelError > ll.level {
		return
	}
	msg := fmt.Sprintf("[E] "+format, v...)
	ll.l.Printf(msg)
}
