package service

import (
	"kafkago/internal/service/request"
)

type Services struct {
	Request *request.Service
}

func InitServices(request *request.Service) *Services {
	return &Services{
		Request: request,
	}
}
