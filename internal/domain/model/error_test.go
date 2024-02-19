package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusError(t *testing.T) {
	tests := []struct {
		name             string
		givenStatusError *StatusError
		expectedRes      string
	}{
		{
			"unexpected error",
			&StatusError{},
			"error: unexpected error",
		},
		{
			"status error",
			ErrBadRequest,
			"error: bad request",
		},
		{
			"nil",
			nil,
			"nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.givenStatusError.Error()
			assert.Equal(t, tt.expectedRes, res)
		})
	}
}
