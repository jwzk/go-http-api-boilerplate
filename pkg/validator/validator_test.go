package validator

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestValidator_Validate(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		type dto struct {
			Name string `json:"name" validate:"required"`
		}

		jsonDTO, err := json.Marshal(dto{Name: "Name"})
		assert.NoError(t, err)

		err = Validate[dto](context.Background(), bytes.NewBuffer(jsonDTO), &dto{})
		assert.NoError(t, err)
	})

	t.Run("error decode body", func(t *testing.T) {
		type dto struct {
			Title string `json:"title" validate:"required"`
		}
		var errTarget *json.SyntaxError

		err := Validate[dto](context.Background(), strings.NewReader("body"), &dto{})
		assert.ErrorAs(t, err, &errTarget)
	})

	t.Run("error struct validation", func(t *testing.T) {
		type dto struct {
			Title string `json:"title" validate:"required"`
		}
		var errTarget validator.ValidationErrors

		jsonDTO, err := json.Marshal(dto{})
		assert.NoError(t, err)

		err = Validate[dto](context.Background(), bytes.NewBuffer(jsonDTO), &dto{})
		assert.ErrorAs(t, err, &errTarget)
	})
}
