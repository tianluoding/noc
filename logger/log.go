package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func InitLogger() {
	// 初始化 zap 的配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development: true,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "console",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	// 创建 zap.Logger 实例
	var err error
	logger, err := config.Build()
	Logger = logger.Sugar()
	if err != nil {
		fallbackLogger, _ := zap.NewDevelopment()
		Logger = fallbackLogger.Sugar()
	}
}

// type Field zap.Field

// func Infof(format string, v ...interface{}) {
// 	logger.Sugar().Infof(format, v...)
// }

// func Errorf(format string, v ...interface{}) {
// 	logger.Sugar().Errorf(format, v...)
// }

// func Fatalf(format string, v ...interface{}) {
// 	logger.Sugar().Fatalf(format, v...)
// }

// func Panicf(format string, v ...interface{}) {
// 	logger.Sugar().Panicf(format, v...)
// }

// func Info(v ...interface{}) {
// 	logger.Sugar().Info(v...)
// }

// func Error(v ...interface{}) {
// 	logger.Sugar().Error(v...)
// }

// func Fatal(v ...interface{}) {
// 	logger.Sugar().Fatal(v...)
// }

// func Panic(v ...interface{}) {
// 	logger.Sugar().Panic(v...)
// }
