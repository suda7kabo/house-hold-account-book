package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Logger *zap.Logger
}

func NewLogger() (*Logger, error) {
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	zapLogger, err := config.Build()
	if err != nil {
		return nil, err
	}
	return &Logger{Logger: zapLogger}, nil
}

func (l Logger) Sync() error {
	return l.Logger.Sync()
}

func (l Logger) Fatal(msg string, err error) {
	l.Logger.Fatal(msg, zap.Error(err))
}

func (l Logger) Error(msg string, err error) {
	l.Logger.Error(msg, zap.Error(err))
}
func (l Logger) Info(msg string, err error) {
	l.Logger.Info(msg, zap.Error(err))
}
