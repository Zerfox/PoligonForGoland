package REST_API

import (
	"Poligon/REST API/http"
	"Poligon/REST API/todo"
	"fmt"
)

func MainTodo() {
	todoList := todo.NewList()
	httpHandlers := http.NewHTTPHandlers(todoList)
	httpServer := http.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server", err)
	}
}
