package logger

import (
	"encoding/json"
)

func GetAppLogger() Logger {
	return appLoggerConfig
}

func (log LogConfig) Info(message string) {
	log.logger.Info().Msg(message)
}

func (log LogConfig) InfoInterface(data interface{}) {
	marshaledData, _ := json.Marshal(data)
	log.logger.Info().Msg(string(marshaledData))
}

func (log LogConfig) Error(message string) {
	log.logger.Error().Msg(message)
}

func (log LogConfig) Debug(message string) {
	log.logger.Debug().Msg(message)
}
