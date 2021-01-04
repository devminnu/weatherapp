package logger

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger *zap.Logger

// Should be called only from main!
func Init() {
	var err error
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stderr", "logs/log.log"}
	cfg.ErrorOutputPaths = []string{"stderr", "logs/log.log"}

	logger, err = cfg.Build()
	if err != nil {
		log.Fatalf("Error while setting logger")
		return
	}

	gin.SetMode("debug")

	logger.Warn("Warning")
	logger.Error("Error")
	logger.Info("Info")
}

func Get() *zap.Logger {
	return logger
}

func Close() {
	logger.Sync()
}

// Any other custom configuration for logging or output
