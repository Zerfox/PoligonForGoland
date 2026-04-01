package BodyHTTP

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserHandler struct {
	Name        string `json:"Name"`
	Address     string `json:"Address"`
	Age         int    `json:"Age"`
	IsHeMarried string `json:"IsHeMarried"`
	Height      int    `json:"Height"`
}

func JSONmain() {
	http.HandleFunc("/UserHandler", testJSON)

	fmt.Println("Start server")
	err := http.ListenAndServe(":9098", nil)
	if err != nil {
		fmt.Println("Произошла ошибка в  работе сервера", err.Error())
	}

	fmt.Println("Stop server")
}

func (c UserHandler) Println() {
	fmt.Println("\nName", c.Name)
	fmt.Println("Address", c.Address)
	fmt.Println("Age", c.Age)
	fmt.Println("IsHeMarried", c.IsHeMarried)
	fmt.Println("Height", c.Height)
}
func testJSON(w http.ResponseWriter, r *http.Request) {

	var userHandler UserHandler
	if err := json.NewDecoder(r.Body).Decode(&userHandler); err != nil {
		fmt.Println("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	httpWritter := UserHandler{
		Name:        "Ля_Пердоле",
		Address:     "Бобр це бидло",
		Age:         5,
		IsHeMarried: "no",
		Height:      133}

	b, err := json.Marshal(httpWritter)
	if err != nil {
		fmt.Println("Не поолучилось обработать ответ на запрос")
		w.WriteHeader(http.StatusInsufficientStorage)
		return
	}

	if _, err = w.Write(b); err != nil {
		fmt.Println("Не поолучилось отправить ответ на запрос", err)
		w.WriteHeader(http.StatusInsufficientStorage)
		return
	}

	userHandler.Println()
}
