package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {
	// 初始化 zap 的配置
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Development:      false,
		Sampling:         &zap.SamplingConfig{},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	// 创建 zap.Logger 实例
	var err error
	Logger, err = config.Build()
	if err != nil {
		panic(err)
	}
}

type Field zap.Field

func Infof(format string, v ...interface{}) {
	Logger.Sugar().Infof(format, v...)
}

func Errorf(format string, v ...interface{}) {
	Logger.Sugar().Errorf(format, v...)
}

func Fatalf(format string, v ...interface{}) {
	Logger.Sugar().Fatalf(format, v...)
}

func Panicf(format string, v ...interface{}) {
	Logger.Sugar().Panicf(format, v...)
}

func Info(v ...interface{}) {
	Logger.Sugar().Info(v...)
}

func Error(v ...interface{}) {
	Logger.Sugar().Error(v...)
}

func Fatal(v ...interface{}) {
	Logger.Sugar().Fatal(v...)
}

func Panic(v ...interface{}) {
	Logger.Sugar().Panic(v...)
}
