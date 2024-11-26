package kafka

import (
	"crypto/tls"
	"crypto/x509"

	"kafkago/internal/config"

	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func sasl(cfg config.Kafka) (*kafka.Transport, error) {
	if !cfg.Sasl.Enabled {
		return &kafka.Transport{}, nil
	}
	certs := x509.NewCertPool()
	certs.AppendCertsFromPEM([]byte(cfg.Sasl.Cert))
	saslMechanism, err := scram.Mechanism(scram.SHA512, cfg.Sasl.User, cfg.Sasl.Password)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create SASL mechanism")
	}
	return &kafka.Transport{
		TLS: &tls.Config{
			RootCAs: certs,
			// nolint gosec
			InsecureSkipVerify: true,
		},
		SASL: saslMechanism,
	}, nil
}
