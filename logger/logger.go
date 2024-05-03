package logger

import "go.uber.org/zap"

var Logger *zap.Logger

func init() {
	cfg := zap.NewDevelopmentConfig()
	cfg.DisableStacktrace = true
	l, err := cfg.Build()
	if err != nil {
		panic("Failed to create logger: " + err.Error())
	}
	Logger = l
}

func GetLogger() *zap.Logger {
	return Logger
}
