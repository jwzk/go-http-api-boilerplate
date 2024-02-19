package book

import (
	"github.com/go-chi/chi/v5"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/port"
)

type BookRouter struct {
	l port.Logger

	getBookUC    port.GetBook
	getBooksUC   port.GetBooks
	createBookUC port.CreateBook
	updateBookUC port.UpdateBook
	deleteBookUC port.DeleteBook
}

func NewBookRouter(
	l port.Logger,

	getBookUC port.GetBook,
	getBooksUC port.GetBooks,
	createBookUC port.CreateBook,
	updateBookUC port.UpdateBook,
	deleteBookUC port.DeleteBook,
) chi.Router {
	var (
		b = &BookRouter{
			l,
			getBookUC, getBooksUC, createBookUC, updateBookUC, deleteBookUC,
		}
		r = chi.NewRouter()
	)

	r.Get("/", b.getBooks())
	r.Post("/", b.createBook())
	r.Get("/{bookID}", b.getBook())
	r.Put("/{bookID}", b.updateBook())
	r.Delete("/{bookID}", b.deleteBook())

	return r
}
