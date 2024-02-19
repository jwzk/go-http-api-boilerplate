package book

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
)

func TestBook_Model(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		dto := bookDTO{Title: "Title", Author: "Author", PublicationDate: time.Unix(1, 1)}
		dtoModel := dto.Model()

		expectedModel := model.Book{Title: "Title", Author: "Author", PublicationDate: time.Unix(1, 1)}

		assert.Equal(t, expectedModel, dtoModel)
	})
}
