package validator

import (
	"context"
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)

func Validate[T any](ctx context.Context, r io.Reader, dto *T) error {
	err := json.NewDecoder(r).Decode(dto)
	if err != nil {
		return err
	}

	validate := validator.New()

	err = validate.StructCtx(ctx, dto)
	if err != nil {
		return err
	}

	return nil
}
