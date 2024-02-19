package memory

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
)

func TestBookDAO_CRUD(t *testing.T) {
	givenBook := model.Book{ID: "1", Title: "Title", Author: "Author", PublicationDate: time.Unix(1, 1)}

	t.Run("CRUD ok", func(t *testing.T) {
		ctx := context.Background()
		dao := NewBookDAO()

		savedBook, err := dao.CreateBook(ctx, givenBook)
		assert.NoError(t, err)
		assert.Equal(t, givenBook.Title, savedBook.Title)
		assert.Equal(t, givenBook.Author, savedBook.Author)
		assert.Equal(t, givenBook.PublicationDate, savedBook.PublicationDate)

		gettedBook, err := dao.GetBook(ctx, savedBook.ID)
		assert.NoError(t, err)
		assert.Equal(t, savedBook, gettedBook)

		gettedBooks, err := dao.GetBooks(ctx)
		assert.NoError(t, err)
		assert.Equal(t, []model.Book{savedBook}, gettedBooks)

		givenUpdateBook := savedBook
		givenUpdateBook.Title = "Updated Title"
		givenUpdateBook.Author = "Updated Author"
		givenUpdateBook.PublicationDate = time.Unix(2, 2)

		updatedBook, err := dao.UpdateBook(ctx, givenUpdateBook)
		assert.NoError(t, err)
		assert.Equal(t, givenUpdateBook, updatedBook)
		assert.NotEqual(t, savedBook, updatedBook)

		err = dao.DeleteBook(ctx, savedBook.ID)
		assert.NoError(t, err)

		_, err = dao.GetBook(ctx, savedBook.ID)
		assert.Equal(t, model.ErrBookNotFound, err)
	})

	t.Run("GetBook not found", func(t *testing.T) {
		ctx := context.Background()
		dao := NewBookDAO()

		_, err := dao.GetBook(ctx, "1")
		assert.Equal(t, model.ErrBookNotFound, err)
	})

	t.Run("UpdateBook not found", func(t *testing.T) {
		ctx := context.Background()
		dao := NewBookDAO()

		_, err := dao.UpdateBook(ctx, givenBook)
		assert.Equal(t, model.ErrBookNotFound, err)
	})

	t.Run("DeleteBook not found", func(t *testing.T) {
		ctx := context.Background()
		dao := NewBookDAO()

		err := dao.DeleteBook(ctx, "1")
		assert.Equal(t, model.ErrBookNotFound, err)
	})
}

func TestBookDAO_generateID(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		dao := NewBookDAO()

		id := dao.generateID()
		assert.Equal(t, 10, len(id))
	})
}
