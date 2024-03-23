package book_inventory_system_logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

const (
	outputFormat = "stdout"
)

type Option func(loggerConfig *zap.Config) error

func InitConfig() *zap.Config {
	return &zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      nil,
		ErrorOutputPaths: nil,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "time",
			LevelKey:      "level",
			NameKey:       "logger",
			CallerKey:     "caller",
			MessageKey:    "message",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
	}
}

func WithConsoleOutput() Option {
	return func(lc *zap.Config) error {
		if lc.OutputPaths != nil {
			lc.OutputPaths = append(lc.OutputPaths, outputFormat)
		} else {
			lc.OutputPaths = []string{outputFormat}
		}

		return nil
	}
}

func WithFileOutput(filepath string) Option {
	return func(lc *zap.Config) error {
		if err := createLogsDirectory(filepath); err != nil {
			return fmt.Errorf("failed to create log directory: %v", err)
		}

		if lc.OutputPaths != nil {
			lc.OutputPaths = append(lc.OutputPaths, filepath)
		} else {
			lc.OutputPaths = []string{filepath}
		}

		return nil
	}
}

func createLogsDirectory(loggerOutputFilePath string) error {
	splitPath := strings.Split(loggerOutputFilePath, "/")
	loggerOutputDir := strings.Join(splitPath[:len(splitPath)-1], "/")

	if _, err := os.Stat(loggerOutputDir); os.IsNotExist(err) {
		if err = os.MkdirAll(loggerOutputDir, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}
