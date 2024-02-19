package book

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/jwzk/go-http-api-boilerplate/internal/adapter/http/internal/handler"
	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
	"github.com/jwzk/go-http-api-boilerplate/pkg/validator"
)

func (b *BookRouter) getBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookID := model.BookID(chi.URLParam(r, "bookID"))

		res, err := b.getBookUC(r.Context(), bookID)
		if err != nil {
			handler.JsonResponse(b.l, w, nil, err)

			return
		}

		handler.JsonResponse(b.l, w, res, nil)
	}
}

func (b BookRouter) getBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := b.getBooksUC(r.Context())
		if err != nil {
			handler.JsonResponse(b.l, w, nil, err)

			return
		}

		handler.JsonResponse(b.l, w, books, nil)
	}
}

func (b BookRouter) createBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto bookDTO

		err := validator.Validate[bookDTO](r.Context(), r.Body, &dto)
		if err != nil {
			handler.JsonResponse(b.l, w, nil, model.ErrBadRequest)

			return
		}

		savedBook, err := b.createBookUC(r.Context(), dto.Model())
		if err != nil {
			handler.JsonResponse(b.l, w, nil, err)

			return
		}

		handler.JsonResponse(b.l, w, savedBook, nil)
	}
}

func (b BookRouter) updateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto bookDTO
		bookID := model.BookID(chi.URLParam(r, "bookID"))

		err := validator.Validate[bookDTO](r.Context(), r.Body, &dto)
		if err != nil {
			handler.JsonResponse(b.l, w, nil, model.ErrBadRequest)

			return
		}

		inputBook := dto.Model()
		inputBook.ID = bookID

		savedBook, err := b.updateBookUC(r.Context(), inputBook)
		if err != nil {
			handler.JsonResponse(b.l, w, nil, err)

			return
		}

		handler.JsonResponse(b.l, w, savedBook, nil)
	}
}

func (b BookRouter) deleteBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookID := model.BookID(chi.URLParam(r, "bookID"))

		err := b.deleteBookUC(r.Context(), bookID)
		if err != nil {
			handler.JsonResponse(b.l, w, nil, err)

			return
		}

		handler.JsonResponse(b.l, w, nil, nil)
	}
}
