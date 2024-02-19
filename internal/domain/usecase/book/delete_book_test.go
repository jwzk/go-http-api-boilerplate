package book

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
	"github.com/jwzk/go-http-api-boilerplate/internal/domain/port/mocks"
	"github.com/jwzk/go-http-api-boilerplate/pkg/logger"
)

func TestDeleteBook(t *testing.T) {
	l := logger.New("test")

	tests := []struct {
		name        string
		givenBookID model.BookID
		daoErr      error
		expectedErr error
	}{
		{
			"error",
			model.BookID("1"),
			errors.New("test error"),
			fmt.Errorf("book dao: %w", errors.New("test error")),
		},
		{
			"ok",
			model.BookID("1"),
			nil,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockBookDAO := mocks.NewMockBookDAO(t)

			mockBookDAO.
				EXPECT().
				DeleteBook(ctx, tt.givenBookID).
				Return(tt.daoErr).
				Once()

			uc := NewDeleteBook(l, mockBookDAO)

			err := uc(ctx, tt.givenBookID)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
