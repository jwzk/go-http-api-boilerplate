package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ValidateDTO[K any](r *http.Request, dto *K) error {
	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		return err
	}

	validate := validator.New()

	err := validate.StructCtx(r.Context(), dto)
	if err != nil {
		return err
	}

	return nil
}
