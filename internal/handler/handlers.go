package handler

import (
	"kafkago/internal/config"
	"kafkago/internal/handler/request"
	"kafkago/internal/service"
	"kafkago/pkg/httputils"
)

type Handlers struct {
	validator *httputils.CustomValidator
	services  *service.Services
	request   *request.Handler
}

func InitHandlers(
	_ *config.Config,
	services *service.Services,
	validator *httputils.CustomValidator,
) *Handlers {
	return &Handlers{
		validator: validator,
		services:  services,
		request:   request.NewHandler(services.Request, validator),
	}
}
