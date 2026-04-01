package Postman

import (
	"fmt"
	"net/http"
)

func dogHandler(w http.ResponseWriter, r *http.Request) {
	str := []byte("Я собака и я говорю Гав")

	_, err := w.Write(str)
	if err != nil {
		fmt.Println("Во время выполнения запроса произошла ошибка", err)
	} else {
		fmt.Println("Успешная обработка собакенского запроса")
	}
}
func catHandler(w http.ResponseWriter, r *http.Request) {
	str := []byte("Я кошка и я говорю Мяу ")

	_, err := w.Write(str)
	if err != nil {
		fmt.Println("Во время выполнения запроса произошла ошибка", err)
	} else {
		fmt.Println("Успешная обработка кошачего запроса")
	}
}
func humanHandler(w http.ResponseWriter, r *http.Request) {

	_, err := w.Write([]byte("Я человек и я говорю Java скоро умрет"))
	if err != nil {
		fmt.Println("Во время выполнения запроса произошла ошибка", err)
	} else {
		fmt.Println("Успешная обработка хуманского запроса")
	}
}

func FirstContactMain() {
	http.HandleFunc("/dog", dogHandler)
	http.HandleFunc("/cat", catHandler)
	http.HandleFunc("/human", humanHandler)

	fmt.Println("Start server")
	err := http.ListenAndServe(":9098", nil)
	if err != nil {
		fmt.Println("Произошла ошибка в  работе сервера", err.Error())
	}

	fmt.Println("Stop server")
}
