package model

import (
	"fmt"
	"time"
)

var ErrBookNotFound = fmt.Errorf("book %w", ErrNotFound)

type BookID string

type Book struct {
	ID              BookID    `json:"id"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	PublicationDate time.Time `json:"publicationDate"`
}
