package rlog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

type ctxLoggerKey struct{}

func InitLog(handlers ...slog.Handler) {
	var h slog.Handler
	if len(handlers) == 0 {
		h = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			//AddSource: true,
			Level: slog.LevelInfo,
		})
	} else {
		h = slog.NewMultiHandler(handlers...)
	}
	l := slog.New(h)
	slog.SetDefault(l)
}

func WithCxt(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, ctxLoggerKey{}, logger)
}

func FromCxt(ctx context.Context) *slog.Logger {
	if log, ok := ctx.Value(ctxLoggerKey{}).(*slog.Logger); ok {
		return log
	}
	return slog.Default()
}

func With(ctx context.Context, args ...any) context.Context {
	subLogger := FromCxt(ctx).With(args...)
	return WithCxt(ctx, subLogger)
}

func InfoF(format string, args ...any) {
	slog.Info(fmt.Sprintf(format, args...))
}

func Info(content string, attrs ...any) {
	slog.Info(content, attrs...)
}

func InfoCxt(cxt context.Context, content string, attrs ...any) {
	slog.InfoContext(cxt, content, attrs...)
}

func InfoFCxt(cxt context.Context, format string, args ...any) {
	slog.InfoContext(cxt, fmt.Sprintf(format, args...))
}

func Errorf(format string, args ...any) {
	slog.Error(fmt.Sprintf(format, args...))
}
