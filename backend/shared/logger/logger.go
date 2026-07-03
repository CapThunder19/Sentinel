package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func Init() {
	Logger, _ = zap.NewDevelopment()
}

func Info(message string) {
	Logger.Info(message)

}

func Error(message string) {
	Logger.Error(message)

}

func Warn(message string) {
	Logger.Warn(message)

}

func Debug(message string){
	Logger.Debug(message)

}