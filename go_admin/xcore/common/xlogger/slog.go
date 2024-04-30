package xlogger

import (
	"context"
	"golang.org/x/exp/slog"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

type Logger struct {
	infoLogger  *slog.Logger
	errorLogger *slog.Logger
	logPath     string
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
		MaxSize:    1,
		MaxAge:     30,
		MaxBackups: 10,
		Compress:   true,
	}
	var _logger Logger
	_logger.logPath = logPath
	//info 打印
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

	_logger.infoLogger = slog.New(slog.NewJSONHandler(infoLumberjack, &_infoOptions))
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
	_logger.errorLogger = slog.New(slog.NewJSONHandler(errLumberjack, &_errOptions))
	//startCheckFoldCron(&_logger)
	return &_logger

}

func (receiver *Logger) Error(ctx context.Context, msg string, args ...any) {
	receiver.infoLogger.ErrorContext(ctx, msg, args)
}

func (receiver *Logger) Info(msg string, args ...any) {
	receiver.errorLogger.Info(msg, args)
}
