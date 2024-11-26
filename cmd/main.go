package main

import (
	"kafkago/internal/app"
	"kafkago/internal/bootstrap"
	"kafkago/internal/broker"
	"kafkago/internal/config"
	"kafkago/internal/handler"
	"kafkago/internal/service"
	"kafkago/internal/service/request"
	"kafkago/pkg/httputils"
	"kafkago/pkg/kafka"
	"log"
	"log/slog"
)

func main() {
	cfg, err := config.CreateConfig()
	if err != nil {
		log.Fatalf("failed to create config: %v", err)
	}

	bootstrap.InitLogger(cfg.Logger)
	logger := slog.Default()

	app := app.New(cfg.AppConfig, logger)

	ctx := app.GetShutdownContext()

	// Initialize Kafka
	kafkaDialer, err := kafka.New(ctx, cfg.Kafka)
	if err != nil {
		log.Fatalf("failed to initialize Kafka: %v", err)
	}
	defer kafkaDialer.Close()

	// Initialize brokers
	brokers, err := broker.InitBrokers(kafkaDialer)
	if err != nil {
		log.Fatalf("failed to initialize brokers: %v", err)
	}

	// Initialize services
	services := service.InitServices(
		request.NewService(
			cfg,
			brokers.Producer,
		),
	)

	validator, err := httputils.NewValidator()
	if err != nil {
		log.Fatalf("failed to initialize validator: %v", err)
	}

	handlers := handler.InitHandlers(cfg, services, validator)

	handler.InitRouter(app, handlers)
	app.Run()
}
