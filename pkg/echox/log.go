package echox

import (
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

type Logger struct {
	*log.Logger
	zapLogger zap.Logger
}

func NewLogger(zapLogger zap.Logger) *Logger {
	return &Logger{Logger: log.New(""), zapLogger: zapLogger}
}

