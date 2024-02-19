package bookapi

import (
	"net/http"

	pkghttp "github.com/jwzk/go-http-api-boilerplate/pkg/http"
)

func NewRouter(
	bookRouter http.Handler,
) *http.ServeMux {
	apiRouter := http.NewServeMux()

	pkghttp.Mount(apiRouter, "/books", bookRouter)

	return apiRouter
}
