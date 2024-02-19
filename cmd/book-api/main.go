package main

import (
	"flag"

	bookmemorydao "github.com/jwzk/go-http-api-boilerplate/internal/adapter/dao/book/memory"
	"github.com/jwzk/go-http-api-boilerplate/internal/adapter/http/bookapi"
	"github.com/jwzk/go-http-api-boilerplate/internal/adapter/http/bookapi/book"
	bookusecase "github.com/jwzk/go-http-api-boilerplate/internal/domain/usecase/book"
	"github.com/jwzk/go-http-api-boilerplate/pkg/logger"
)

func main() {
	// Flag
	loggerLevel := flag.String("level", "info", "logger level (debug|info)")
	httpPort := flag.String("port", "4100", "http api port")
	flag.Parse()

	// Logger
	l := logger.New(*loggerLevel)

	// DAO
	bookDAO := bookmemorydao.NewBookDAO()

	// Usecase
	getBookUC := bookusecase.NewGetBook(l, bookDAO)
	getBooksUC := bookusecase.NewGetBooks(l, bookDAO)
	createBookUC := bookusecase.NewCreateBook(l, bookDAO)
	updateBookUC := bookusecase.NewUpdateBook(l, bookDAO)
	deleteBookUC := bookusecase.NewDeleteBook(l, bookDAO)

	// Router
	bookRouter := book.NewBookRouter(l, getBookUC, getBooksUC, createBookUC, updateBookUC, deleteBookUC)

	// HTTP Server
	s := bookapi.NewHTTPServer(l, *httpPort, l.GetMiddleware(), bookRouter)
	l.Infof("listening and serving HTTP on :%s", *httpPort)

	// Run
	l.Panic(s.ListenAndServe())
}
