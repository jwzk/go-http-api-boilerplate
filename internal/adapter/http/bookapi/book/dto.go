package book

import (
	"time"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
)

type bookDTO struct {
	Title           string    `json:"title"           validate:"required"`
	Author          string    `json:"author"          validate:"required"`
	PublicationDate time.Time `json:"publicationDate" validate:"required"`
}

func (dto bookDTO) Model() model.Book {
	return model.Book{
		Title:           dto.Title,
		Author:          dto.Author,
		PublicationDate: dto.PublicationDate,
	}
}
