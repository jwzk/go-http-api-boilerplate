package book

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
	"github.com/jwzk/go-http-api-boilerplate/pkg/logger"
)

func TestBook_NewBookRouter(t *testing.T) {
	l := logger.New("test")
	givenBook := model.Book{ID: "1", Title: "Title", Author: "Author", PublicationDate: time.Unix(1, 1)}

	jsonDTO, err := json.Marshal(bookDTO{Title: "Title", Author: "Author", PublicationDate: time.Unix(1, 1)})
	assert.NoError(t, err)

	r := NewBookRouter(
		l,
		func(ctx context.Context, bookID model.BookID) (model.Book, error) {
			return givenBook, nil
		},
		func(ctx context.Context) ([]model.Book, error) {
			return []model.Book{givenBook}, nil
		},
		func(ctx context.Context, book model.Book) (model.Book, error) {
			return givenBook, nil
		},
		func(ctx context.Context, book model.Book) (model.Book, error) {
			return givenBook, nil
		},
		func(ctx context.Context, bookID model.BookID) error {
			return nil
		},
	)

	t.Run("ok GET /", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})

	t.Run("ok POST /{bookID}", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonDTO))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})

	t.Run("ok GET /{bookID}", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/1", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})

	t.Run("ok PUT /{bookID}", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/1", bytes.NewBuffer(jsonDTO))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})

	t.Run("ok DELETE /{bookID}", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/1", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Result().StatusCode)
	})
}
