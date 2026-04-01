package http

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHTTPServer(handlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{httpHandlers: handlers}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/books").Methods("POST").HandlerFunc(s.httpHandlers.HandlerAddBook)
	router.Path("/books/{name}").Methods("GET").HandlerFunc(s.httpHandlers.HandlerGetBook)
	router.Path("/books").Methods("GET").Queries("completed", "true").HandlerFunc(s.httpHandlers.HandlerGetAllUncompletedBook)
	router.Path("/books").Methods("GET").HandlerFunc(s.httpHandlers.HandlerGetAllBook)
	router.Path("/books/{name}").Methods("PATCH").HandlerFunc(s.httpHandlers.HandlerCompleteBook)
	router.Path("/books/{name}").Methods("DELETE").HandlerFunc(s.httpHandlers.HandlerDeliteBook)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
