package http

import (
	"net/http"
)

// Mount securely mounts a sub-router at a prefix.
// It supports exact prefix matches without the default trailing slash redirects
func Mount(mux *http.ServeMux, prefix string, handler http.Handler) {
	mux.Handle(prefix+"/", http.StripPrefix(prefix, handler))

	mux.Handle(prefix, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r2 := r.Clone(r.Context())
		r2.URL.Path = "/"
		handler.ServeHTTP(w, r2)
	}))
}
