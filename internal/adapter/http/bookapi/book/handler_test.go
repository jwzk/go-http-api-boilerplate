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

	"github.com/stretchr/testify/assert"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
)

func TestBook_getBook(t *testing.T) {
	givenBook := model.Book{ID: "1", Title: "Title", Author: "Author", PublicationDate: time.Unix(1, 1).UTC()}

	t.Run("ok 200", func(t *testing.T) {
		uc := func(ctx context.Context, bookID model.BookID) (model.Book, error) {
			return givenBook, nil
		}

		mux := http.NewServeMux()
		bookRouter := BookRouter{getBookUC: uc}
		mux.HandleFunc("GET /{bookID}", bookRouter.getBook())

		req := httptest.NewRequest(http.MethodGet, "/1", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

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

		mux := http.NewServeMux()
		bookRouter := BookRouter{getBookUC: uc}
		mux.HandleFunc("GET /{bookID}", bookRouter.getBook())

		req := httptest.NewRequest(http.MethodGet, "/1234", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})

	t.Run("error dao 500", func(t *testing.T) {
		uc := func(ctx context.Context, bookID model.BookID) (model.Book, error) {
			return model.Book{}, errors.New("test error")
		}

		mux := http.NewServeMux()
		bookRouter := BookRouter{getBookUC: uc}
		mux.HandleFunc("GET /{bookID}", bookRouter.getBook())

		req := httptest.NewRequest(http.MethodGet, "/1234", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}

func TestBook_getBooks(t *testing.T) {
	givenBook := model.Book{ID: "1", Title: "Title", Author: "Author", PublicationDate: time.Unix(1, 1).UTC()}

	t.Run("ok 200", func(t *testing.T) {
		uc := func(ctx context.Context) ([]model.Book, error) {
			return []model.Book{givenBook}, nil
		}

		mux := http.NewServeMux()
		bookRouter := &BookRouter{getBooksUC: uc}
		mux.HandleFunc("GET /{$}", bookRouter.getBooks())

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

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

		mux := http.NewServeMux()
		bookRouter := &BookRouter{getBooksUC: uc}
		mux.HandleFunc("GET /{$}", bookRouter.getBooks())

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}

func TestBook_createBook(t *testing.T) {
	givenBook := model.Book{ID: "1", Title: "Title", Author: "Author", PublicationDate: time.Unix(1, 1).UTC()}

	jsonDTO, err := json.Marshal(bookDTO{Title: "Title", Author: "Author", PublicationDate: time.Unix(1, 1).UTC()})
	assert.NoError(t, err)

	t.Run("ok 200", func(t *testing.T) {
		uc := func(ctx context.Context, book model.Book) (model.Book, error) {
			return givenBook, nil
		}

		mux := http.NewServeMux()
		bookRouter := &BookRouter{createBookUC: uc}
		mux.HandleFunc("POST /{$}", bookRouter.createBook())

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonDTO))
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)

		var res model.Book
		err := json.NewDecoder(w.Result().Body).Decode(&res)
		assert.NoError(t, err)

		assert.Equal(t, givenBook, res)
	})

	t.Run("error dto validation bad request 404", func(t *testing.T) {
		mux := http.NewServeMux()
		bookRouter := &BookRouter{}
		mux.HandleFunc("POST /{$}", bookRouter.createBook())

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("{}")))
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("error dao 500", func(t *testing.T) {
		uc := func(ctx context.Context, book model.Book) (model.Book, error) {
			return model.Book{}, errors.New("test error")
		}

		mux := http.NewServeMux()
		bookRouter := &BookRouter{createBookUC: uc}
		mux.HandleFunc("POST /{$}", bookRouter.createBook())

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonDTO))
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}

func TestBook_updateBook(t *testing.T) {
	givenBook := model.Book{ID: "1", Title: "Title", Author: "Author", PublicationDate: time.Unix(1, 1).UTC()}

	jsonDTO, err := json.Marshal(bookDTO{Title: "Title", Author: "Author", PublicationDate: time.Unix(1, 1).UTC()})
	assert.NoError(t, err)

	t.Run("ok 200", func(t *testing.T) {
		uc := func(ctx context.Context, book model.Book) (model.Book, error) {
			return givenBook, nil
		}

		mux := http.NewServeMux()
		bookRouter := &BookRouter{updateBookUC: uc}
		mux.HandleFunc("PUT /{bookID}", bookRouter.updateBook())

		req := httptest.NewRequest(http.MethodPut, "/1", bytes.NewBuffer(jsonDTO))
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)

		var res model.Book
		err := json.NewDecoder(w.Result().Body).Decode(&res)
		assert.NoError(t, err)

		assert.Equal(t, givenBook, res)
	})

	t.Run("error dto validation bad request 404", func(t *testing.T) {
		mux := http.NewServeMux()
		bookRouter := &BookRouter{}
		mux.HandleFunc("PUT /{bookID}", bookRouter.updateBook())

		req := httptest.NewRequest(http.MethodPut, "/1", bytes.NewBuffer([]byte("{}")))
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("error dao not found 404", func(t *testing.T) {
		uc := func(ctx context.Context, book model.Book) (model.Book, error) {
			return model.Book{}, model.ErrNotFound
		}

		mux := http.NewServeMux()
		bookRouter := &BookRouter{updateBookUC: uc}
		mux.HandleFunc("PUT /{bookID}", bookRouter.updateBook())

		req := httptest.NewRequest(http.MethodPut, "/1234", bytes.NewBuffer(jsonDTO))
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})

	t.Run("error dao 500", func(t *testing.T) {
		uc := func(ctx context.Context, book model.Book) (model.Book, error) {
			return model.Book{}, errors.New("test error")
		}

		mux := http.NewServeMux()
		bookRouter := &BookRouter{updateBookUC: uc}
		mux.HandleFunc("PUT /{bookID}", bookRouter.updateBook())

		req := httptest.NewRequest(http.MethodPut, "/1", bytes.NewBuffer(jsonDTO))
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}

func TestBook_deleteBook(t *testing.T) {
	t.Run("ok 200", func(t *testing.T) {
		uc := func(ctx context.Context, bookID model.BookID) error {
			return nil
		}

		mux := http.NewServeMux()
		bookRouter := &BookRouter{deleteBookUC: uc}
		mux.HandleFunc("DELETE /{bookID}", bookRouter.deleteBook())

		req := httptest.NewRequest(http.MethodDelete, "/1", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Result().StatusCode)
	})

	t.Run("error dao not found 404", func(t *testing.T) {
		uc := func(ctx context.Context, bookID model.BookID) error {
			return model.ErrNotFound
		}

		mux := http.NewServeMux()
		bookRouter := &BookRouter{deleteBookUC: uc}
		mux.HandleFunc("DELETE /{bookID}", bookRouter.deleteBook())

		req := httptest.NewRequest(http.MethodDelete, "/1234", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})

	t.Run("error dao 500", func(t *testing.T) {
		uc := func(ctx context.Context, bookID model.BookID) error {
			return errors.New("test error")
		}

		mux := http.NewServeMux()
		bookRouter := &BookRouter{deleteBookUC: uc}
		mux.HandleFunc("DELETE /{bookID}", bookRouter.deleteBook())

		req := httptest.NewRequest(http.MethodDelete, "/1", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}
