package request

import (
	"kafkago/internal/config"
	"log/slog"

	broker "kafkago/internal/broker/request"
)

type Service struct {
	config   *config.Config
	logger   *slog.Logger
	producer *broker.Producer
}

func NewService(
	cfg *config.Config,
	producer *broker.Producer,
) *Service {
	return &Service{
		config:   cfg,
		logger:   slog.Default(),
		producer: producer,
	}
}
