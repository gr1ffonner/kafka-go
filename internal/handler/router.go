package handler

import (
	"kafkago/internal/app"
	"kafkago/pkg/middleware"
)

func InitRouter(frame *app.App, handlers *Handlers) {
	// Request
	{
		frame.RegisterHTTPHandler(
			app.Get, "/kafka",
			middleware.Logging(
				handlers.request.KafkaWrite,
				"request",
			),
		)
	}
}
