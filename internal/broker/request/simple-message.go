package request

import (
	"context"
	"kafkago/internal/broker/request/models"

	"github.com/pkg/errors"
)

func (p *Producer) SendSimpleMessage(ctx context.Context, topic string, message models.Msg) error {
	err := p.dialer.WriteWithRetry(ctx, topic, message, 3)
	if err != nil {
		return errors.Wrap(err, "failed to send message")
	}

	return nil
}
