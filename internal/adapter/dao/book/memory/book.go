package memory

import (
	"context"
	"math/rand"
	"sync"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type bookDAO struct {
	books map[model.BookID]model.Book
	mutex sync.RWMutex
}

func NewBookDAO() *bookDAO {
	return &bookDAO{
		books: make(map[model.BookID]model.Book),
	}
}

func (dao *bookDAO) GetBook(ctx context.Context, bookID model.BookID) (model.Book, error) {
	dao.mutex.RLock()
	defer dao.mutex.RUnlock()

	book, ok := dao.books[bookID]
	if !ok {
		return model.Book{}, model.ErrNotFound
	}

	return book, nil
}

func (dao *bookDAO) GetBooks(ctx context.Context) ([]model.Book, error) {
	dao.mutex.RLock()
	defer dao.mutex.RUnlock()

	var books []model.Book
	for _, b := range dao.books {
		books = append(books, b)
	}

	return books, nil
}

func (dao *bookDAO) CreateBook(ctx context.Context, book model.Book) (model.Book, error) {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()

	book.ID = dao.generateID()
	dao.books[book.ID] = book

	return book, nil
}

func (dao *bookDAO) UpdateBook(ctx context.Context, book model.Book) (model.Book, error) {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()

	_, ok := dao.books[book.ID]
	if !ok {
		return model.Book{}, model.ErrNotFound
	}

	dao.books[book.ID] = book

	return book, nil
}

func (dao *bookDAO) DeleteBook(ctx context.Context, bookID model.BookID) error {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()

	_, ok := dao.books[bookID]
	if !ok {
		return model.ErrNotFound
	}

	delete(dao.books, bookID)

	return nil
}

func (dao *bookDAO) generateID() model.BookID {
	randomString := make([]byte, 10)
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}

	return model.BookID(randomString)
}
