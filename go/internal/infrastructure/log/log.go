package log

import "go.uber.org/zap"

func GetLogger() *zap.SugaredLogger {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	sugar := logger.Sugar()

	return sugar
}
