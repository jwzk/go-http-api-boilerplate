package book

import (
	"context"
	"fmt"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
	"github.com/jwzk/go-http-api-boilerplate/internal/domain/port"
)

type deleteBook struct {
	dao port.BookDAO
}

func NewDeleteBook(dao port.BookDAO) port.DeleteBook {
	uc := &deleteBook{dao}
	return uc.deleteBook
}

func (uc *deleteBook) deleteBook(ctx context.Context, bookID model.BookID) error {
	err := uc.dao.DeleteBook(ctx, bookID)
	if err != nil {
		return fmt.Errorf("book dao: %w", err)
	}

	return nil
}
