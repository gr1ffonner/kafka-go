package kafka

import (
	"fmt"
	"log/slog"
)

type (
	errLogger slog.Logger
	logger    slog.Logger
)

func (l errLogger) Printf(format string, v ...interface{}) {
	(*slog.Logger)(&l).Error("kafka", "msg", fmt.Sprintf(format, v...))
}

func (l logger) Printf(format string, v ...interface{}) {
	(*slog.Logger)(&l).Info("kafka", "msg", fmt.Sprintf(format, v...))
}
