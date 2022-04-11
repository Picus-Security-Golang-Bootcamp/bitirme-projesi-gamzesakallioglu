package logger

import (
	"fmt"

	"github.com/gamze.sakallioglu/learningGo/bitirme-projesi-gamzesakallioglu/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(config *config.Config) {
	logLevel, err := zapcore.ParseLevel(config.Logger.Level)
	if err != nil {
		panic(fmt.Sprintf("Unknown log level %v", logLevel))
	}

	//var cfg zap.Config
	cfg := zap.NewDevelopmentConfig() // Development mode

	logger, err := cfg.Build()
	if err != nil {
		logger = zap.NewNop()
	}
	zap.ReplaceGlobals(logger) // for able to access logger globally

}

func Close() {
	defer zap.L().Sync()
}
