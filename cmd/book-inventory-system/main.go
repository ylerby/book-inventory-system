package main

import (
	config "book-inventory-system/internal/config"
	handler "book-inventory-system/internal/handler"
	repository "book-inventory-system/internal/repository"
	service "book-inventory-system/internal/service"
	logger "book-inventory-system/pkg/logger"
	"context"
	"flag"
	"go.uber.org/zap"
	"log"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	outputFileLogPath := flag.String("logfilepath", "", "output log filepath")
	cfgFilePath := flag.String("cfgfilepath", "", "cfg file path")
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	var (
		l    logger.Logger
		err  error
		once sync.Once
	)

	once.Do(func() {
		if *outputFileLogPath != "" {
			l, err = logger.New(
				logger.WithConsoleOutput(),
				logger.WithFileOutput(*outputFileLogPath),
			)
		} else {
			l, err = logger.New(
				logger.WithConsoleOutput(),
			)
		}

		if err != nil {
			log.Fatalf("error occured: %v", err)
		}
	})

	l.Info("init logger")

	cfg, err := config.New(*cfgFilePath)
	if err != nil {
		l.Fatalf("failed to initialize config: %v", err)
	}

	l.Info("init config")

	r, err := repository.New(
		repository.WithDump(
			cfg.Admins,
			cfg.Authors,
			cfg.Books,
			cfg.Genres,
			cfg.Instances,
			cfg.Languages,
			cfg.Productions,
			cfg.Readers,
			cfg.Users,
		),
	)

	if err != nil {
		l.Fatal("failed to initialize repository: %v", err)
	}

	l.Info("init dump/service")

	s := service.New(
		r,
		l.With(zap.String("component", "service")),
		cfg,
	)

	l.Info("init service")

	h := handler.New(
		l,
		s,
	)
	if err != nil {
		l.Fatalf("failed to initialize server: %v", err)
	}

	l.Info("init handlers")

	errCh := make(chan error, 1)
	go h.InitRoutes(
		cfg.ServerAddress,
		errCh,
	)

	l.Info("init routing")

	for {
		select {
		case err := <-errCh:
			l.Fatalf("failed to initialize router: %v", err)

		case <-ctx.Done():
			l.Info("context cancelling")
			return
		}
	}
}
