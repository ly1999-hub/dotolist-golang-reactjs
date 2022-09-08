package logger

import (
	"go.uber.org/zap"
)

type (
	// LogData for tracking
	LogData map[string]interface{}
)

var (
	zapLogger *zap.Logger
)
