package book

import (
	"net/http"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/port"
)

type BookRouter struct {
	getBookUC    port.GetBook
	getBooksUC   port.GetBooks
	createBookUC port.CreateBook
	updateBookUC port.UpdateBook
	deleteBookUC port.DeleteBook
}

func NewBookRouter(
	getBookUC port.GetBook,
	getBooksUC port.GetBooks,
	createBookUC port.CreateBook,
	updateBookUC port.UpdateBook,
	deleteBookUC port.DeleteBook,
) *http.ServeMux {
	var (
		b = &BookRouter{
			getBookUC, getBooksUC, createBookUC, updateBookUC, deleteBookUC,
		}
		mux = http.NewServeMux()
	)

	mux.HandleFunc("GET /{$}", b.getBooks())
	mux.HandleFunc("POST /{$}", b.createBook())
	mux.HandleFunc("GET /{bookID}", b.getBook())
	mux.HandleFunc("PUT /{bookID}", b.updateBook())
	mux.HandleFunc("DELETE /{bookID}", b.deleteBook())

	return mux
}
