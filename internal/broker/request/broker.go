package request

import (
	"kafkago/pkg/kafka"
	"log/slog"
)

type Producer struct {
	dialer *kafka.Dialer
	l      *slog.Logger
}

func NewProducer(dialer *kafka.Dialer) *Producer {
	return &Producer{
		dialer: dialer,
		l:      slog.Default(),
	}
}
