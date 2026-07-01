package logging

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"

	"github.com/faith/color"
)

type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type PrettyHandler struct{
	slog.HandlerOptions
	l *log.logger
}

func (h * PrettyHandler) Handle(ctx contect.Context, r slog.Record) error{
	level := r.Level.String() + ":"
	Switch r.Level{
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]any, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool{
		fields[a.Key] = a.Value.Any()
		return true
	})

	timeStr := r.Time.Format("[2006-01-02 15:04:05.000]")
	msg := color.CyanString(r.Message)

	// Only print fields if there are any
	
}