package http_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	pkghttp "github.com/jwzk/go-http-api-boilerplate/pkg/http"
)

func TestMount(t *testing.T) {
	// Arrange sub-router
	subMux := http.NewServeMux()
	subMux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("root"))
	})
	subMux.HandleFunc("GET /{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprintf("id:%s", r.PathValue("id"))))
	})

	// Arrange main router
	mainMux := http.NewServeMux()
	pkghttp.Mount(mainMux, "/books", subMux)

	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "exact match without trailing slash",
			method:         http.MethodGet,
			path:           "/books",
			expectedStatus: http.StatusOK,
			expectedBody:   "root",
		},
		{
			name:           "sub-path match",
			method:         http.MethodGet,
			path:           "/books/123",
			expectedStatus: http.StatusOK,
			expectedBody:   "id:123",
		},
		{
			name:           "not found on main router",
			method:         http.MethodGet,
			path:           "/other",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "404 page not found\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			req := httptest.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()
			mainMux.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, tt.expectedStatus, w.Result().StatusCode)
			assert.Equal(t, tt.expectedBody, w.Body.String())
		})
	}
}
