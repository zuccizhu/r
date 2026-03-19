package rlog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

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
