package helper

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func InitLogger() {
	var core zapcore.Core

	// encoder config
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	if os.Getenv("APP_ENV") == "production" {
		// production - JSON format, save to file
		fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
		logFile, _ := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
		writer := zapcore.AddSync(logFile)
		core = zapcore.NewCore(fileEncoder, writer, zapcore.InfoLevel)
	} else {
		// development - console format
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		core = zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)

	}
	Log = zap.New(core,zap.AddCaller())
}
