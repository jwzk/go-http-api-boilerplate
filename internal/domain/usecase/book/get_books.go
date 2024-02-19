package book

import (
	"context"
	"fmt"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
	"github.com/jwzk/go-http-api-boilerplate/internal/domain/port"
)

type getBooks struct {
	dao port.BookDAO
}

func NewGetBooks(dao port.BookDAO) port.GetBooks {
	uc := &getBooks{dao}
	return uc.getBooks
}

func (uc *getBooks) getBooks(ctx context.Context) ([]model.Book, error) {
	books, err := uc.dao.GetBooks(ctx)
	if err != nil {
		return nil, fmt.Errorf("book dao: %w", err)
	}

	return books, nil
}
