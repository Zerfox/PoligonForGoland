package http

import (
	"Poligon/REST API/todo"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func NewHTTPHandlers(todoList *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
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

func (h *HTTPHandlers) HandlerCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	if err := taskDTO.ValidateForCreate(); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
	}
	todoTask := todo.NewTask(taskDTO.Title, taskDTO.Description)
	if err := h.todoList.AddTask(todoTask); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, todo.ErrTaskAlreadyExistst) {
			http.Error(w, errDTO.ToString(), http.StatusConflict)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}

	b, err := json.MarshalIndent(todoTask, "", "	")
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

func (h *HTTPHandlers) HandlerGetTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	task, err := h.todoList.GetTask(title)
	if err != nil {
		errDTO := ErrorDTO{ // нужно вывести в конструктор
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}
	b, err := json.MarshalIndent(task, "", "		")
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

func (h *HTTPHandlers) HandlerGetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.todoList.ListTask()
	b, err := json.MarshalIndent(tasks, "", "    ")
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

func (h *HTTPHandlers) HandlerGetAllUncompletedTasks(w http.ResponseWriter, r *http.Request) {
	uncompletedTasks := h.todoList.ListUncompletedTask()
	b, err := json.MarshalIndent(uncompletedTasks, "", "    ")
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

func (h *HTTPHandlers) HandlerCompletedTask(w http.ResponseWriter, r *http.Request) {
	var completeDTO CompletedTaskDTO
	if err := json.NewDecoder(r.Body).Decode(&completeDTO); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	title := mux.Vars(r)["title"]

	var (
		changetTask todo.Task
		err         error
	)
	if completeDTO.Complete {
		changetTask, err = h.todoList.CompletedTask(title)
	} else {
		changetTask, err = h.todoList.UncompletedTask(title)
	}
	if err != nil {
		errDTO := ErrorDTO{ // нужно вывести в конструктор
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}

	b, err := json.MarshalIndent(changetTask, "", "    ")
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response: ", err)
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

func (h *HTTPHandlers) HandlerDeliteTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	if err := h.todoList.DeleteTask(title); err != nil {
		errDTO := ErrorDTO{ // нужно вывести в конструктор
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
