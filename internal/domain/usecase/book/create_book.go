package book

import (
	"context"
	"fmt"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
	"github.com/jwzk/go-http-api-boilerplate/internal/domain/port"
)

type createBook struct {
	dao port.BookDAO
}

func NewCreateBook(dao port.BookDAO) port.CreateBook {
	uc := &createBook{dao}
	return uc.createBook
}

func (uc *createBook) createBook(ctx context.Context, book model.Book) (model.Book, error) {
	b, err := uc.dao.CreateBook(ctx, book)
	if err != nil {
		return b, fmt.Errorf("book dao: %w", err)
	}

	return b, nil
}
