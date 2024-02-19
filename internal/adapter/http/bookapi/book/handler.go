package book

import (
	"net/http"

	"github.com/jwzk/go-http-api-boilerplate/internal/adapter/http/internal/writer"
	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
	"github.com/jwzk/go-http-api-boilerplate/pkg/validator"
)

func (b *BookRouter) getBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookID := model.BookID(r.PathValue("bookID"))

		res, err := b.getBookUC(r.Context(), bookID)
		if err != nil {
			writer.JSON(r.Context(), w, nil, err)

			return
		}

		writer.JSON(r.Context(), w, res, nil)
	}
}

func (b BookRouter) getBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := b.getBooksUC(r.Context())
		if err != nil {
			writer.JSON(r.Context(), w, nil, err)

			return
		}

		writer.JSON(r.Context(), w, books, nil)
	}
}

func (b BookRouter) createBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto bookDTO

		err := validator.Validate[bookDTO](r.Context(), r.Body, &dto)
		if err != nil {
			writer.JSON(r.Context(), w, nil, model.ErrBadRequest)

			return
		}

		savedBook, err := b.createBookUC(r.Context(), dto.Model())
		if err != nil {
			writer.JSON(r.Context(), w, nil, err)

			return
		}

		writer.JSON(r.Context(), w, savedBook, nil)
	}
}

func (b BookRouter) updateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto bookDTO
		bookID := model.BookID(r.PathValue("bookID"))

		err := validator.Validate[bookDTO](r.Context(), r.Body, &dto)
		if err != nil {
			writer.JSON(r.Context(), w, nil, model.ErrBadRequest)

			return
		}

		inputBook := dto.Model()
		inputBook.ID = bookID

		savedBook, err := b.updateBookUC(r.Context(), inputBook)
		if err != nil {
			writer.JSON(r.Context(), w, nil, err)

			return
		}

		writer.JSON(r.Context(), w, savedBook, nil)
	}
}

func (b BookRouter) deleteBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookID := model.BookID(r.PathValue("bookID"))

		err := b.deleteBookUC(r.Context(), bookID)
		if err != nil {
			writer.JSON(r.Context(), w, nil, err)

			return
		}

		writer.JSON(r.Context(), w, nil, nil)
	}
}
