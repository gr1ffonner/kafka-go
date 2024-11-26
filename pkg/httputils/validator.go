package httputils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"kafkago/internal/domain/domainerrors"
)

type CustomValidator struct {
	valid *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.valid.Struct(i); err != nil {
		return fmt.Errorf("error while validating data | %w", err)
	}

	return nil
}

func (cv *CustomValidator) DecodeAndValidate(r *http.Request, dst any) error {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(domainerrors.ErrBadRequest, "failed to read request body")
	}
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	err = json.NewDecoder(bytes.NewBuffer(bodyBytes)).Decode(dst)
	if err != nil {
		var syntaxErr *json.SyntaxError
		if errors.As(err, &syntaxErr) {
			return errors.Wrapf(domainerrors.ErrBadRequest, "syntax error in JSON: offset %d, error %v", syntaxErr.Offset, syntaxErr.Error())
		}

		var typeErr *json.UnmarshalTypeError
		if errors.As(err, &typeErr) {
			return errors.Wrapf(domainerrors.ErrBadRequest, "type mismatch in JSON: expected %v, got %v at offset %d", typeErr.Type, typeErr.Value, typeErr.Offset)
		}

		return errors.Wrap(domainerrors.ErrBadRequest, "invalid JSON")
	}

	err = cv.Validate(dst)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errMsg string
			for _, fieldError := range validationErrors {
				errMsg += fmt.Sprintf("You should provide %s: %s; ", fieldError.Field(), fieldError.Error())
			}
			return errors.New(errMsg)
		}
		return errors.Wrap(domainerrors.ErrBadRequest, "JSON validation failed")
	}

	return nil
}

func NewValidator() (*CustomValidator, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return &CustomValidator{valid: validate}, nil
}
