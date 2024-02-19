package port

import (
	"context"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
)

type BookDAO interface {
	GetBook(ctx context.Context, bookID model.BookID) (model.Book, error)
	GetBooks(ctx context.Context) ([]model.Book, error)
	CreateBook(ctx context.Context, book model.Book) (model.Book, error)
	UpdateBook(ctx context.Context, book model.Book) (model.Book, error)
	DeleteBook(ctx context.Context, bookID model.BookID) error
}
