package logger

import (
	"github.com/uma-co82/shupple2-api/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var x *zap.Logger

func init() {
	var err error
	x, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
}

func Configure(c *config.Logger) {
	encoder := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	encoding := "console"
	if c.JSON {
		encoding = "json"
	}

	lc := zap.Config{
		Level:            c.Level,
		Development:      false,
		Encoding:         encoding,
		EncoderConfig:    encoder,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := lc.Build()
	if err != nil {
		panic(err)
	}

	x = logger
}

func Debug(msg string, fields ...zap.Field) {
	x.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	x.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	x.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	x.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	x.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	x.Fatal(msg, fields...)
}
