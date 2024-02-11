package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func MustNewLogger() *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	return logger.Sugar()
}
