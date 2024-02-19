package book

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
	"github.com/jwzk/go-http-api-boilerplate/internal/domain/port/mocks"
	"github.com/jwzk/go-http-api-boilerplate/pkg/logger"
)

func TestGetBooks(t *testing.T) {
	l := logger.New("test")

	tests := []struct {
		name        string
		daoRes      []model.Book
		daoErr      error
		expectedRes []model.Book
		expectedErr error
	}{
		{
			"error",
			nil,
			errors.New("test error"),
			nil,
			fmt.Errorf("book dao: %w", errors.New("test error")),
		},
		{
			"ok",
			[]model.Book{{ID: "1", Title: "Title", Author: "Author", PublicationDate: time.Unix(1, 1)}},
			nil,
			[]model.Book{{ID: "1", Title: "Title", Author: "Author", PublicationDate: time.Unix(1, 1)}},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockBookDAO := mocks.NewMockBookDAO(t)

			mockBookDAO.
				EXPECT().
				GetBooks(ctx).
				Return(tt.daoRes, tt.daoErr).
				Once()

			uc := NewGetBooks(l, mockBookDAO)

			res, err := uc(ctx)
			assert.Equal(t, tt.expectedRes, res)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
