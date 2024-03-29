package loggerPkg

import (
	"os"
	"strings"
)

import (
	"fmt"
	"log"
)

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarning
	LevelError
)

type Logger struct {
	level LogLevel
}

func NewLogger(level LogLevel) *Logger {
	return &Logger{level: level}
}

func (l *Logger) Debug(args ...interface{}) {
	message := fmt.Sprintln(args...)
	message = strings.TrimSuffix(message, "\n")
	l.log(LevelDebug, message)
}

func (l *Logger) Info(message string) {
	l.log(LevelInfo, message)
}

func (l *Logger) Warning(message string) {
	l.log(LevelWarning, message)
}

func (l *Logger) Errorf(format string, err error) error {
	message := fmt.Sprintf(format, " - ", err.Error())
	l.log(LevelError, message)
	return err
}

func (l *Logger) Error(message string) {
	l.log(LevelError, message)
}

func (l *Logger) log(level LogLevel, message string) {
	if level >= l.level {
		levelString := getLevelString(level)
		logMessage := fmt.Sprintf("[%s] %s", levelString, message)
		log.Println(logMessage)
	}
}

func getLevelString(level LogLevel) string {
	switch level {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarning:
		return "WARNING"
	case LevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

func init() {
	log.SetOutput(os.Stdout)
}
