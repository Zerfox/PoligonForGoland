package BodyHTTP

import (
	"fmt"
	"net/http"
	"strconv"
)

func QueryMain() {
	http.HandleFunc("/easeHandler", easeHandler)

	fmt.Println("Start server")
	err := http.ListenAndServe(":9098", nil)
	if err != nil {
		fmt.Println("Произошла ошибка в  работе сервера", err.Error())
	}

}
func easeHandler(w http.ResponseWriter, r *http.Request) {
	oneParam := r.URL.Query().Get("one")
	twoParam := r.URL.Query().Get("two")

	oneP, err := strconv.Atoi(oneParam)
	if err != nil {
		fmt.Println("Не удалось распарсить int")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	twoP, err := strconv.ParseBool(twoParam)
	if err != nil {
		fmt.Println("Не удалось распарсить bool")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("OneParam: ", oneParam)
	fmt.Println("OneP: ", oneP)

	fmt.Println("twoParam: ", twoParam)
	fmt.Println("twoP: ", twoP)
}
