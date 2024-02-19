package bookapi

import (
	"github.com/go-chi/chi/v5"
)

func NewRouter(
	bookRouter chi.Router,
) *chi.Mux {
	apiRouter := chi.NewRouter()
	apiRouter.Mount("/books", bookRouter)

	return apiRouter
}
