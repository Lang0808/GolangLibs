package log

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLog() error {
	var fileInfo *os.File
	var fileError *os.File
	var err error

	logInfoDir := viper.GetString("log.info.dir")
	logErrorDir := viper.GetString("log.error.dir")

	if logInfoDir == "" || logErrorDir == "" {
		return errors.New("log.info.dir or log.error.dir is not specified in config")
	}

	fileInfo, err = os.OpenFile(logInfoDir, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	fileError, err = os.OpenFile(logErrorDir, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	encoderCfg := zapcore.EncoderConfig{
		MessageKey: "Message",
		TimeKey:    "Time",
	}
	encoder := zapcore.NewJSONEncoder(encoderCfg)

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.Lock(fileInfo), lowPriority),
		zapcore.NewCore(encoder, zapcore.Lock(fileError), highPriority),
	)

	logger = zap.New(core)
	return nil
}

func Error(message string) {
	defer logger.Sync()
	logger.Error(message)
}

func Entry(logEntry LogEntry) {
	defer logger.Sync()
	bytes, err := json.Marshal(logEntry)
	if err != nil {
		// do nothing
	}
	logger.Info(string(bytes))
}

func Info(message string) {
	defer logger.Sync()
	logger.Error(message)
}
