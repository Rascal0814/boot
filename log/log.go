package log

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.com/greyxor/slogor"
	"log/slog"
	"os"
	"time"
)

type Logger struct {
	logger *slog.Logger
	name   string
	ctx    context.Context
}

type ServiceName string

const (
	msgk = "msg"
)

func NewLogger(name ServiceName) *Logger {
	logger := slog.New(slogor.NewHandler(os.Stderr, &slogor.Options{
		TimeFormat: time.Stamp,
		Level:      slog.LevelDebug,
		ShowSource: false,
	}))
	return &Logger{logger: logger, name: string(name)}
}

func (l *Logger) Error(msg string, err error, args ...any) error {
	if err != nil {
		l.logger.ErrorContext(l.ctx, l.name, append([]any{msgk, msg, slogor.Err(err)}, args...)...)
		return errors.Wrap(err, msg)
	}

	return nil
}

// Info logs at LevelInfo.
func (l *Logger) Info(msg string, args ...any) {
	l.logger.InfoContext(l.ctx, l.name, append([]any{msgk, msg}, args...)...)
}

// Warn logs at LevelWarn.
func (l *Logger) Warn(msg string, args ...any) {
	l.logger.WarnContext(l.ctx, l.name, append([]any{msgk, msg}, args...)...)
}

// Debug logs at LevelDebug.
func (l *Logger) Debug(msg string, args ...any) {
	l.logger.DebugContext(l.ctx, l.name, append([]any{msgk, msg}, args...)...)
}
