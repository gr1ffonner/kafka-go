package request

import (
	"kafkago/internal/app"
	"kafkago/internal/service/request"
	"kafkago/pkg/httputils"
	"net/http"

	jsoniter "github.com/json-iterator/go"

	"github.com/pkg/errors"
)

type Status struct {
	Status string `json:"status"`
}

type Handler struct {
	s *request.Service
	v *httputils.CustomValidator
}

func NewHandler(service *request.Service, validator *httputils.CustomValidator) *Handler {
	return &Handler{
		s: service,
		v: validator,
	}
}

func (h *Handler) KafkaWrite(r *http.Request) (response *app.HTTPResponse, err error) {
	err = h.s.KafkaWrite(r.Context())
	if err != nil {
		return httputils.FinalizeResponse(nil, err)
	}

	status := Status{"message sended to kafka"}
	v, err := jsoniter.Marshal(status)
	if err != nil {
		return nil, errors.Wrap(err, "marshal status")
	}
	return httputils.FinalizeResponse(&app.HTTPResponse{
		Data: v,
		Code: 200,
	}, nil)
}
