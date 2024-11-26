package httputils

import (
	"kafkago/internal/app"
	"kafkago/internal/domain/domainerrors"
	"net/http"

	"github.com/pkg/errors"
)

type CommonResponse struct {
	Data  any          `json:"data"`
	Error *CommonError `json:"error"`
}

type CommonError struct {
	Code    uint16         `json:"code"`
	Message string         `json:"message"`
	Details map[string]any `json:"details"`
}

func FinalizeResponse(response *app.HTTPResponse, err error) (*app.HTTPResponse, error) {
	if response == nil {
		response = &app.HTTPResponse{}
	}
	response.Headers.Set("Content-Type", "application/json; charset=utf-8")
	response.Headers.Set("Cache-Control", "no-cache, no-store, must-revalidate")
	response.Headers.Set("Pragma", "no-cache")
	response.Headers.Set("Expires", "0")

	if err != nil {
		switch {
		case errors.Is(err, domainerrors.ErrBadRequest):
			response.Code = http.StatusBadRequest
		case errors.Is(err, domainerrors.ErrNotFound):
			response.Code = http.StatusNotFound
		case errors.Is(err, domainerrors.ErrAlreadyExists):
			response.Code = http.StatusNoContent
		default:
			response.Code = http.StatusInternalServerError
		}
	} else {
		response.Code = http.StatusOK
	}

	return response, err
}
