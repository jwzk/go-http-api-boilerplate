package book

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
	"github.com/jwzk/go-http-api-boilerplate/pkg/logger"
)

func TestBook_getBook(t *testing.T) {
	l := logger.New("test")
	givenBook := model.Book{ID: "1", Title: "Title", Author: "Author", PublicationDate: time.Unix(1, 1)}

	t.Run("ok 200", func(t *testing.T) {
		uc := func(ctx context.Context, bookID model.BookID) (model.Book, error) {
			return givenBook, nil
		}

		r := chi.NewRouter()
		bookRouter := BookRouter{l: l, getBookUC: uc}
		r.Get("/{id}", bookRouter.getBook())

		req := httptest.NewRequest(http.MethodGet, "/1", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)

		var res model.Book
		err := json.NewDecoder(w.Result().Body).Decode(&res)
		assert.NoError(t, err)

		assert.Equal(t, givenBook, res)
	})

	t.Run("error dao not found 404", func(t *testing.T) {
		uc := func(ctx context.Context, bookID model.BookID) (model.Book, error) {
			return model.Book{}, model.ErrNotFound
		}

		r := chi.NewRouter()
		bookRouter := BookRouter{l: l, getBookUC: uc}
		r.Get("/{id}", bookRouter.getBook())

		req := httptest.NewRequest(http.MethodGet, "/1234", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})

	t.Run("error dao 500", func(t *testing.T) {
		uc := func(ctx context.Context, bookID model.BookID) (model.Book, error) {
			return model.Book{}, errors.New("test error")
		}

		r := chi.NewRouter()
		bookRouter := BookRouter{l: l, getBookUC: uc}
		r.Get("/{id}", bookRouter.getBook())

		req := httptest.NewRequest(http.MethodGet, "/1234", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}

func TestBook_getBooks(t *testing.T) {
	l := logger.New("test")
	givenBook := model.Book{ID: "1", Title: "Title", Author: "Author", PublicationDate: time.Unix(1, 1)}

	t.Run("ok 200", func(t *testing.T) {
		uc := func(ctx context.Context) ([]model.Book, error) {
			return []model.Book{givenBook}, nil
		}

		r := chi.NewRouter()
		bookRouter := &BookRouter{l: l, getBooksUC: uc}
		r.Get("/", bookRouter.getBooks())

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)

		var res []model.Book
		err := json.NewDecoder(w.Result().Body).Decode(&res)
		assert.NoError(t, err)

		assert.Equal(t, []model.Book{givenBook}, res)
	})

	t.Run("error dao 500", func(t *testing.T) {
		uc := func(ctx context.Context) ([]model.Book, error) {
			return []model.Book{}, errors.New("test error")
		}

		r := chi.NewRouter()
		bookRouter := &BookRouter{l: l, getBooksUC: uc}
		r.Get("/", bookRouter.getBooks())

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}

func TestBook_createBook(t *testing.T) {
	l := logger.New("test")
	givenBook := model.Book{ID: "1", Title: "Title", Author: "Author", PublicationDate: time.Unix(1, 1)}

	jsonDTO, err := json.Marshal(bookDTO{Title: "Title", Author: "Author", PublicationDate: time.Unix(1, 1)})
	assert.NoError(t, err)

	t.Run("ok 200", func(t *testing.T) {
		uc := func(ctx context.Context, book model.Book) (model.Book, error) {
			return givenBook, nil
		}

		r := chi.NewRouter()
		bookRouter := &BookRouter{l: l, createBookUC: uc}
		r.Post("/", bookRouter.createBook())

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonDTO))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)

		var res model.Book
		err := json.NewDecoder(w.Result().Body).Decode(&res)
		assert.NoError(t, err)

		assert.Equal(t, givenBook, res)
	})

	t.Run("error dto validation bad request 404", func(t *testing.T) {
		r := chi.NewRouter()
		bookRouter := &BookRouter{l: l}
		r.Post("/", bookRouter.createBook())

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("{}")))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("error dao 500", func(t *testing.T) {
		uc := func(ctx context.Context, book model.Book) (model.Book, error) {
			return model.Book{}, errors.New("test error")
		}

		r := chi.NewRouter()
		bookRouter := &BookRouter{l: l, createBookUC: uc}
		r.Post("/", bookRouter.createBook())

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonDTO))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}

func TestBook_updateBook(t *testing.T) {
	l := logger.New("test")
	givenBook := model.Book{ID: "1", Title: "Title", Author: "Author", PublicationDate: time.Unix(1, 1)}

	jsonDTO, err := json.Marshal(bookDTO{Title: "Title", Author: "Author", PublicationDate: time.Unix(1, 1)})
	assert.NoError(t, err)

	t.Run("ok 200", func(t *testing.T) {
		uc := func(ctx context.Context, book model.Book) (model.Book, error) {
			return givenBook, nil
		}

		r := chi.NewRouter()
		bookRouter := &BookRouter{l: l, updateBookUC: uc}
		r.Put("/{id}", bookRouter.updateBook())

		req := httptest.NewRequest(http.MethodPut, "/1", bytes.NewBuffer(jsonDTO))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)

		var res model.Book
		err := json.NewDecoder(w.Result().Body).Decode(&res)
		assert.NoError(t, err)

		assert.Equal(t, givenBook, res)
	})

	t.Run("error dto validation bad request 404", func(t *testing.T) {
		r := chi.NewRouter()
		bookRouter := &BookRouter{l: l}
		r.Put("/{id}", bookRouter.updateBook())

		req := httptest.NewRequest(http.MethodPut, "/1", bytes.NewBuffer([]byte("{}")))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("error dao not found 404", func(t *testing.T) {
		uc := func(ctx context.Context, book model.Book) (model.Book, error) {
			return model.Book{}, model.ErrNotFound
		}

		r := chi.NewRouter()
		bookRouter := &BookRouter{l: l, updateBookUC: uc}
		r.Put("/{id}", bookRouter.updateBook())

		req := httptest.NewRequest(http.MethodPut, "/1234", bytes.NewBuffer(jsonDTO))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})

	t.Run("error dao 500", func(t *testing.T) {
		uc := func(ctx context.Context, book model.Book) (model.Book, error) {
			return model.Book{}, errors.New("test error")
		}

		r := chi.NewRouter()
		bookRouter := &BookRouter{l: l, updateBookUC: uc}
		r.Put("/{id}", bookRouter.updateBook())

		req := httptest.NewRequest(http.MethodPut, "/1", bytes.NewBuffer(jsonDTO))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}

func TestBook_deleteBook(t *testing.T) {
	l := logger.New("test")

	t.Run("ok 200", func(t *testing.T) {
		uc := func(ctx context.Context, bookID model.BookID) error {
			return nil
		}

		r := chi.NewRouter()
		bookRouter := &BookRouter{l: l, deleteBookUC: uc}
		r.Delete("/{id}", bookRouter.deleteBook())

		req := httptest.NewRequest(http.MethodDelete, "/1", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Result().StatusCode)
	})

	t.Run("error dao not found 404", func(t *testing.T) {
		uc := func(ctx context.Context, bookID model.BookID) error {
			return model.ErrNotFound
		}

		r := chi.NewRouter()
		bookRouter := &BookRouter{l: l, deleteBookUC: uc}
		r.Delete("/{id}", bookRouter.deleteBook())

		req := httptest.NewRequest(http.MethodDelete, "/1234", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})

	t.Run("error dao 500", func(t *testing.T) {
		uc := func(ctx context.Context, bookID model.BookID) error {
			return errors.New("test error")
		}

		r := chi.NewRouter()
		bookRouter := &BookRouter{l: l, deleteBookUC: uc}
		r.Delete("/{id}", bookRouter.deleteBook())

		req := httptest.NewRequest(http.MethodDelete, "/1", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}
