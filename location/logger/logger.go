package logger

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger *zap.Logger

// Should be called only from main!
func Init() {
	// gin.DisableConsoleColor()

	// // Logging to a file.
	// f, _ := os.Create("logs/log.json")
	// gin.DefaultWriter = io.MultiWriter(f)
	// zapLogger, err := zap.NewProduction()
	// if err != nil {
	// 	panic(err)
	// }
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

	// logger = zapLogger //.Sugar()
	// zapLogger.Warn("Warning")
	// zapLogger.Error("Error")
	// zapLogger.Info("Info")
}

func Get() *zap.Logger {
	return logger
}

func Close() {
	logger.Sync()
}

// Any other custom configuration for logging or output
