package book

import (
	"context"
	"fmt"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
	"github.com/jwzk/go-http-api-boilerplate/internal/domain/port"
)

type updateBook struct {
	l   port.Logger
	dao port.BookDAO
}

func NewUpdateBook(l port.Logger, dao port.BookDAO) port.UpdateBook {
	uc := &updateBook{l, dao}
	return uc.updateBook
}

func (uc *updateBook) updateBook(ctx context.Context, book model.Book) (model.Book, error) {
	b, err := uc.dao.UpdateBook(ctx, book)
	if err != nil {
		return b, fmt.Errorf("book dao: %w", err)
	}

	return b, nil
}
