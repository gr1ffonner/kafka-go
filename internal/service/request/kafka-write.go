package request

import (
	"context"
	"kafkago/internal/broker/request/models"

	"github.com/pkg/errors"
)

func (s *Service) KafkaWrite(
	ctx context.Context,
) error {
	msg := models.Msg{
		MessageID: "228",
		Message:   "pisya popa",
	}
	err := s.producer.SendSimpleMessage(ctx, s.config.Kafka.Topics.TestTopic, msg)
	if err != nil {
		return errors.Wrap(err, "send simple message")
	}
	return nil
}
