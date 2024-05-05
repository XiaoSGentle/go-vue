package xlogger

import (
	"golang.org/x/exp/slog"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

type Logger struct {
	ErrorLog *slog.Logger
	InfoLog  *slog.Logger
}

func NewSlog(logPath string) *Logger {
	errLumberjack := &lumberjack.Logger{
		Filename:   logPath + "/error.log",
		LocalTime:  true,
		MaxSize:    1,
		MaxAge:     30,
		MaxBackups: 10,
		Compress:   true,
	}

	infoLumberjack := &lumberjack.Logger{
		Filename:   logPath + "/info.log",
		LocalTime:  true,
		MaxSize:    2,
		MaxAge:     30,
		MaxBackups: 10,
		Compress:   true,
	}
	//Error 打印
	_errOptions := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelError,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				if t, ok := a.Value.Any().(time.Time); ok {
					a.Value = slog.StringValue(t.Format(time.DateTime))
				}
			}
			return a
		},
	}

	_infoOptions := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				if t, ok := a.Value.Any().(time.Time); ok {
					a.Value = slog.StringValue(t.Format(time.DateTime))
				}
			}
			return a
		},
	}
	return &Logger{
		ErrorLog: slog.New(slog.NewJSONHandler(errLumberjack, &_errOptions)),
		InfoLog:  slog.New(slog.NewJSONHandler(infoLumberjack, &_infoOptions)),
	}
}
