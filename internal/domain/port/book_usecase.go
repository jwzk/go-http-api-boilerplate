package port

import (
	"context"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
)

type GetBook func(ctx context.Context, bookID model.BookID) (model.Book, error)

type GetBooks func(ctx context.Context) ([]model.Book, error)

type CreateBook func(ctx context.Context, book model.Book) (model.Book, error)

type UpdateBook func(ctx context.Context, book model.Book) (model.Book, error)

type DeleteBook func(ctx context.Context, bookID model.BookID) error
