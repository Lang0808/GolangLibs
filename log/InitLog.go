package log

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/Lang0808/GolangLibs/config"
	"github.com/Lang0808/GolangLibs/file"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLog() error {
	var fileInfo *os.File
	var fileError *os.File
	var err error

	logInfoDir := config.GetString("log.info.dir")
	logErrorDir := config.GetString("log.error.dir")

	if logInfoDir == "" || logErrorDir == "" {
		return errors.New("log.info.dir or log.error.dir is not specified in config")
	}

	fileInfo, err = file.OpenOrCreateFile(logInfoDir)
	if err != nil {
		return err
	}

	fileError, err = file.OpenOrCreateFile(logErrorDir)
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
	logger.Info(message)
}
