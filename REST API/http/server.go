package http

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHTTPServer(HTTPHandler *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: HTTPHandler,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/tasks").Methods("POST").HandlerFunc(s.httpHandlers.HandlerCreateTask)
	router.Path("/tasks/(title)").Methods("GET").HandlerFunc(s.httpHandlers.HandlerGetTask)
	router.Path("/tasks").Methods("GET").Queries("completed", "true").HandlerFunc(s.httpHandlers.HandlerGetAllTasks)
	router.Path("/tasks").Methods("GET").HandlerFunc(s.httpHandlers.HandlerGetAllTasks)
	router.Path("/tasks/(title)").Methods("PATCH").HandlerFunc(s.httpHandlers.HandlerCompletedTask)
	router.Path("/tasks/(title)").Methods("DELETE").HandlerFunc(s.httpHandlers.HandlerDeliteTask)

	if err := http.ListenAndServe("9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
