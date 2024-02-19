package book

import (
	"context"
	"fmt"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
	"github.com/jwzk/go-http-api-boilerplate/internal/domain/port"
)

type getBook struct {
	l   port.Logger
	dao port.BookDAO
}

func NewGetBook(l port.Logger, dao port.BookDAO) port.GetBook {
	uc := &getBook{l, dao}
	return uc.getBook
}

func (uc *getBook) getBook(ctx context.Context, bookID model.BookID) (model.Book, error) {
	b, err := uc.dao.GetBook(ctx, bookID)
	if err != nil {
		return b, fmt.Errorf("book dao: %w", err)
	}

	return b, nil
}
