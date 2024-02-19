package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler_ValidateDTO(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		type dto struct {
			Name string `json:"name" validate:"required"`
		}

		jsonDTO, err := json.Marshal(dto{Name: "Name"})
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(jsonDTO))

		err = ValidateDTO[dto](req, &dto{})
		assert.NoError(t, err)
	})

	t.Run("error decode body", func(t *testing.T) {
		type dto struct {
			Title string `json:"title" validate:"required"`
		}

		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader("body"))

		err := ValidateDTO[dto](req, &dto{})
		assert.Equal(t, "invalid character 'b' looking for beginning of value", err.Error())
	})

	t.Run("error struct validation", func(t *testing.T) {
		type dto struct {
			Title string `json:"title" validate:"required"`
		}

		jsonDTO, err := json.Marshal(dto{})
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(jsonDTO))

		err = ValidateDTO[dto](req, &dto{})
		assert.Equal(t, "Key: 'dto.Title' Error:Field validation for 'Title' failed on the 'required' tag", err.Error())
	})
}
