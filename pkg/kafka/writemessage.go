package kafka

import (
	"context"
	"log/slog"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"

	"github.com/pkg/errors"
)

func (d *Dialer) WriteWithRetry(ctx context.Context, topic string, message any, maxRetries int) error {
	var retryErr error
	for i := 0; i < maxRetries; i++ {
		if err := d.WriteMessage(ctx, topic, message); err != nil {
			retryErr = err
			slog.Default().Error("Failed to write message to Kafka", "attempt", i+1, "error", err)

			if strings.Contains(err.Error(), "Leader Not Available") {
				time.Sleep(time.Second * time.Duration(2<<i))
				continue
			}

			return errors.Wrap(err, "failed to write message to Kafka")
		}
		retryErr = nil
		break
	}

	if retryErr != nil {
		return errors.Wrap(retryErr, "failed to write message to Kafka after retries")
	}

	return nil
}

func (d *Dialer) WriteMessage(ctx context.Context, topic string, data any) error {
	bData, err := jsoniter.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "json marshal error")
	}
	return d.writer.WriteMessages(ctx, kafka.Message{Topic: topic, Value: bData})
}
