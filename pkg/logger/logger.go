package book_inventory_system_logger

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	With(args ...interface{}) *zap.SugaredLogger
	Sync() error
}

type logger struct {
	*zap.SugaredLogger
}

func New(opts ...Option) (Logger, error) {
	errs := make([]error, 0)

	if opts == nil {
		return nil, fmt.Errorf("missing log options")
	}

	config := InitConfig()
	for _, opt := range opts {
		err := opt(config)
		errs = append(errs, err)
	}

	log, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build logger config: %w", err)
	}

	sugar := log.Sugar()

	return &logger{
		SugaredLogger: sugar,
	}, errors.Join(errs...)
}
