package kafka

import (
	"context"
	"kafkago/internal/config"
	"log/slog"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"

	"github.com/pkg/errors"
)

type Dialer struct {
	writer *kafka.Writer
	cfg    config.Kafka
}

func New(ctx context.Context, cfg config.Kafka) (*Dialer, error) {
	secure, err := sasl(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create kafka secure connection")
	}

	dialer := &Dialer{
		writer: &kafka.Writer{
			Logger:                 (logger)(*slog.Default()),
			ErrorLogger:            (errLogger)(*slog.Default()),
			Addr:                   kafka.TCP(strings.Split(cfg.DSN, ",")...),
			RequiredAcks:           kafka.RequireAll,
			Async:                  true,
			WriteTimeout:           time.Duration(cfg.WriteTimeoutSec) * time.Second,
			Transport:              secure,
			AllowAutoTopicCreation: true,
		},
		cfg: cfg,
	}

	err = dialer.warmUpQueue(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "warm up queue")
	}

	return dialer, nil
}

func (d *Dialer) Close() {
	if err := d.writer.Close(); err != nil {
		slog.Default().Error("closing kafka writer", "error", err)
	}
}

func (d *Dialer) warmUpQueue(ctx context.Context) error {
	messages := map[string]string{
		"message": "Kafka dialer successfully initialized.",
	}

	var lastErr error
	for i := 0; i < 5; i++ {
		errState := d.WriteMessage(ctx, d.cfg.Topics.TestTopic, messages)

		if errState == nil {
			return nil
		}
	}

	return lastErr
}
