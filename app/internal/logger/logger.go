package logger

import (
	"github.com/Vityalimbaev/Example-Backend/config"
	"github.com/sirupsen/logrus"
)

func SetupLogger() {
	serverConfig := config.GetServerConfig()
	switch serverConfig.LogLevel {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetReportCaller(serverConfig.LogShowPath)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	default:
		logrus.SetLevel(logrus.ErrorLevel)
	}
}
