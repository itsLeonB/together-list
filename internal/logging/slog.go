package logging

import (
	"fmt"
	"log/slog"
	"os"
)

func Init() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo, // Change to LevelDebug if needed
	}))
	slog.SetDefault(logger)
}

func Debug(args ...any) {
	slog.Debug(fmt.Sprint(args...))
}

func Info(args ...any) {
	slog.Info(fmt.Sprint(args...))
}

func Warn(args ...any) {
	slog.Warn(fmt.Sprint(args...))
}

func Error(args ...any) {
	slog.Error(fmt.Sprint(args...))
}

func Fatal(args ...any) {
	slog.Error(fmt.Sprint(args...))
	os.Exit(1)
}

func Panic(args ...any) {
	msg := fmt.Sprint(args...)
	slog.Error(msg)
	panic(msg)
}

func Debugf(format string, args ...any) {
	slog.Debug(fmt.Sprintf(format, args...))
}

func Infof(format string, args ...any) {
	slog.Info(fmt.Sprintf(format, args...))
}

func Warnf(format string, args ...any) {
	slog.Warn(fmt.Sprintf(format, args...))
}

func Errorf(format string, args ...any) {
	slog.Error(fmt.Sprintf(format, args...))
}

func Fatalf(format string, args ...any) {
	slog.Error(fmt.Sprintf(format, args...))
	os.Exit(1)
}

func Panicf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	slog.Error(msg)
	panic(msg)
}
