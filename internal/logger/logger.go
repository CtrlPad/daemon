package logger

import "github.com/charmbracelet/log"

func Info(msg string, args ...any) {
	log.Info(msg, args...)
}

func Print(msg string, args ...any) {
	log.Print(msg, args...)
}

func Error(msg string, args ...any) {
	log.Error(msg, args...)
}
