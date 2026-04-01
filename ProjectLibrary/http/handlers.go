package http

import (
	"Poligon/ProjectLibrary/library"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type HTTPHandlers struct {
	lib *library.Library
}

func NewHTTPHadlersBook(lib *library.Library) *HTTPHandlers {
	return &HTTPHandlers{
		lib: lib,
	}
}

/*
   Pattern: tasks
   method: POST
   info: JSON in HTTP request body

   succeed:
     - status code : 201 created
     - response body: JSON represent created task

   failed:
     - status code: 400, 409, 500...
     - response body: JSON with error + time
*/

func (h *HTTPHandlers) HandlerAddBook(w http.ResponseWriter, r *http.Request) {
	var bookDTO BookDTO
	if err := json.NewDecoder(r.Body).Decode(&bookDTO); err != nil {
		errDTO := ErrorDTO{Message: err.Error(), Time: time.Now()}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	if err := bookDTO.ValidationForCreate(); err != nil {
		errDTO := ErrorDTO{Message: err.Error(), Time: time.Now()}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
	}
	bookForLibrary := library.NewBook(bookDTO.Name, bookDTO.Autor, bookDTO.NumberOfPages)
	if err := h.lib.AddBook(bookForLibrary); err != nil {
		errDTO := NewErrorDTO(err)
		if errors.Is(err, library.ErrBookAlreadyExist) {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}

	b, err := json.MarshalIndent(bookForLibrary, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return

	}
}

/*
   Pattern: /tasks/(title)
   method: GET
   info: pattern

   succeed:
     - status code : 200 OK
     - response body: JSON represent found task

   failed:
     - status code: 400, 404, 500...
     - response body: JSON with error + time
*/

func (h *HTTPHandlers) HandlerGetBook(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	fmt.Println(name)
	book, err := h.lib.GetBook(name)
	if err != nil {
		errDTO := NewErrorDTO(err)
		if errors.Is(err, library.ErrBookNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}
	b, err := json.MarshalIndent(book, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Failed to write http response", err)
		return
	}

}

/*
   Pattern: /tasks
   method: GET
   info: -

   succeed:
     - status code : 200 OK
     - response body: JSON represent found tasks

   failed:
     - status code: 400, 500...
     - response body: JSON with error + time
*/

func (h *HTTPHandlers) HandlerGetAllBook(w http.ResponseWriter, r *http.Request) {
	books := h.lib.ListBook()
	b, err := json.MarshalIndent(books, "", "     ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Failed to write http response", err)
		return
	}

}

/*
   Pattern: /tasks?completed=true
   method: GET
   info: query params

   succeed:
     - status code : 200 OK
     - response body: JSON represent found tasks

   failed:
     - status code: 400, 500...
     - response body: JSON with error + time
*/

func (h *HTTPHandlers) HandlerGetAllUncompletedBook(w http.ResponseWriter, r *http.Request) {
	uncompletedBook := h.lib.ListUncompletedBook()
	b, err := json.MarshalIndent(uncompletedBook, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Failed to write http response", err)
		return
	}
}

/*
   Pattern: /tasks/(title)
   method: PATCH
   info: pattern + JSON in request body

   succeed:
     - status code : 200 OK
     - response body: JSON represent changed tasks

   failed:
     - status code: 400, 500...
     - response body: JSON with error + time
*/

func (h *HTTPHandlers) HandlerCompleteBook(w http.ResponseWriter, r *http.Request) {
	var completedBookDTO CompletedBookDTO
	fmt.Println("Первично", completedBookDTO)
	if err := json.NewDecoder(r.Body).Decode(&completedBookDTO); err != nil {
		errDTO := NewErrorDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	name := mux.Vars(r)["name"]
	fmt.Println("ля ты ввел", &completedBookDTO)
	var (
		changedBook library.Book
		err         error
	)
	if completedBookDTO.CompleteBook {
		fmt.Println("Прочитали")
		changedBook, err = h.lib.CompletedBook(name)
	} else {
		fmt.Println("читаем заново")
		changedBook, err = h.lib.UncompletedBook(name)
	}
	if err != nil {
		errDTO := NewErrorDTO(err)
		if errors.Is(err, library.ErrBookNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}
	b, err := json.MarshalIndent(changedBook, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Failed to write http response", err)
		return
	}
}

/*
   Pattern: tasks/(title)
   method: DELITE
   info: pattern

   succeed:
     - status code : 204 No Contend
     - response body: -

   failed:
     - status code: 400,404, 500...
     - response body: JSON with error + time
*/

func (h *HTTPHandlers) HandlerDeliteBook(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	if err := h.lib.DeleteBook(name); err != nil {
		errDTO := NewErrorDTO(err)
		if errors.Is(err, library.ErrBookNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}
