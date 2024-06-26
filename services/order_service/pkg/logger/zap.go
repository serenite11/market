package logger_zap

import (
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New() *zap.Logger {
	lg := zap.NewDevelopmentEncoderConfig()
	lg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(lg),
			zapcore.AddSync(colorable.NewColorableStdout()),
			zapcore.DebugLevel,
		),
	)
}
