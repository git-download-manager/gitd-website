package loggers

import (
	"go.uber.org/zap"
)

var Plain *zap.Logger

func SetupPlainLogger(serviceName, serviceVersion string) {
	// General Zap Logger Init
	mainLogger, _ := zap.NewProduction()
	defer mainLogger.Sync()

	Plain = mainLogger
	Plain.Named(serviceName + ":main:" + serviceVersion)
}
