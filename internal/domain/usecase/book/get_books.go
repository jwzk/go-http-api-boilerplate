package book

import (
	"context"
	"fmt"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
	"github.com/jwzk/go-http-api-boilerplate/internal/domain/port"
)

type getBooks struct {
	l   port.Logger
	dao port.BookDAO
}

func NewGetBooks(l port.Logger, dao port.BookDAO) port.GetBooks {
	uc := &getBooks{l, dao}
	return uc.getBooks
}

func (uc *getBooks) getBooks(ctx context.Context) ([]model.Book, error) {
	books, err := uc.dao.GetBooks(ctx)
	if err != nil {
		return nil, fmt.Errorf("book dao: %w", err)
	}

	return books, nil
}
