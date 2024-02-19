package book

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/jwzk/go-http-api-boilerplate/internal/adapter/http/pkg/handler"
	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
)

func (b *BookRouter) getBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookID := model.BookID(chi.URLParam(r, "bookID"))

		book, err := b.getBookUC(r.Context(), bookID)
		if err != nil {
			handler.JsonResponseError(b.l, w, err)
			return
		}

		handler.JsonResponse(w, http.StatusOK, book)
	}
}

func (b BookRouter) getBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := b.getBooksUC(r.Context())
		if err != nil {
			handler.JsonResponseError(b.l, w, err)
			return
		}

		handler.JsonResponse(w, http.StatusOK, books)
	}
}

func (b BookRouter) createBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto bookDTO

		err := handler.ValidateDTO(r, &dto)
		if err != nil {
			handler.JsonResponseError(b.l, w, model.ErrBadRequest)
			return
		}

		savedBook, err := b.createBookUC(r.Context(), dto.Model())
		if err != nil {
			handler.JsonResponseError(b.l, w, err)
			return
		}

		handler.JsonResponse(w, http.StatusCreated, savedBook)
	}
}

func (b BookRouter) updateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto bookDTO
		bookID := model.BookID(chi.URLParam(r, "bookID"))

		err := handler.ValidateDTO(r, &dto)
		if err != nil {
			handler.JsonResponseError(b.l, w, model.ErrBadRequest)
			return
		}

		inputBook := dto.Model()
		inputBook.ID = bookID

		savedBook, err := b.updateBookUC(r.Context(), inputBook)
		if err != nil {
			handler.JsonResponseError(b.l, w, err)
			return
		}

		handler.JsonResponse(w, http.StatusOK, savedBook)
	}
}

func (b BookRouter) deleteBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookID := model.BookID(chi.URLParam(r, "bookID"))

		err := b.deleteBookUC(r.Context(), bookID)
		if err != nil {
			handler.JsonResponseError(b.l, w, err)
			return
		}

		handler.JsonResponse(w, http.StatusOK, nil)
	}
}
