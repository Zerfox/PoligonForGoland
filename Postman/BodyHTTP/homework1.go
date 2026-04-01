package BodyHTTP

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

type Message struct {
	Id          int    `json:"Id"`
	Heading     string `json:"Heading"`
	PostalCode  int    `json:"PostalCode"`
	TextMessage string `json:"TextMessage"`
	Criticality bool   `json:"Criticality"`
}

// Println метод выдает полученные данные в консоль
func (m Message) Println() {
	fmt.Println("\n Id: ", m.Id)
	fmt.Println("Heading:", m.Heading)
	fmt.Println("PostalCode:", m.PostalCode)
	fmt.Println("TextMessage: ", m.TextMessage)
	fmt.Println("Criticality: ", m.Criticality)
}

// Метод получения рандомного числа
func getId() int {
	for i := 0; ; {
		id := rand.Intn(10000)
		// Проверка наличия ключа
		if len(storageMessage) >= 1 {
			if storageMessage[i].Id == id {
				fmt.Println("Дублирование значения ключа, регенирация ключа")
				continue
			} else {
				fmt.Println("Генерация ключа прошла успешно")
				return id
			}
		} else {
			return id
		}

	}
}

var storageMessage = make(map[int]Message)
var mx = sync.Mutex{}

func MessageHandlerMain() {

	http.HandleFunc("/messageSaver", messageSaver)
	http.HandleFunc("/messageLoader", messageLoader)
	http.HandleFunc("/messageRemover", messageRemover)
	http.HandleFunc("/messageSearch", messageSearch)

	fmt.Println("Start server")
	err := http.ListenAndServe(":9098", nil)
	if err != nil {
		fmt.Println("Произошла ошибка в  работе сервера", err.Error())
	}

	fmt.Println("Stop server")
}

// Хендлер сохраняет сообщение в мапу
func messageSaver(w http.ResponseWriter, r *http.Request) {
	var message Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		fmt.Println("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//герация рандомного числа
	message.Id = getId()
	// внесение в мапу сообщения
	mx.Lock()
	storageMessage[message.Id] = message
	mx.Unlock()

	// ответное сообщение в формате JSON
	replyMessage := message
	rM, err := json.Marshal(replyMessage)
	if err != nil {
		fmt.Println("Не получилось обработать ответ на запрос")
		w.WriteHeader(http.StatusInsufficientStorage)
		return
	}
	if _, err = w.Write(rM); err != nil {
		fmt.Println("Не получилось отправить ответ на запрос", err)
		w.WriteHeader(http.StatusInsufficientStorage)
		return
	}

	message.Println()

}

// Хендлер удаляет сообщение из структуры
func messageRemover(w http.ResponseWriter, r *http.Request) {
	idMessage, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Не удалось обработать запрос", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	numberId, err := strconv.Atoi(string(idMessage))
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Println("Не удалось обработать запрос", err)
		return
	}
	_, ok := storageMessage[numberId]
	if ok {
		fmt.Println("Значение найдено")
		_, err = w.Write([]byte("Удаление выполнено успешно"))
		if err != nil {
			fmt.Println("Ошибка отправки ответа на запрос")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		mx.Lock()
		delete(storageMessage, numberId)
		mx.Unlock()
		messageLoader(w, r)
	} else {
		fmt.Println("Не обнаружено совпадений")
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

// Хендлер выполняет вывод сохраненных сообщений
func messageLoader(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("\n Список сохраненных значений"))
	if err != nil {
		fmt.Println("Ошибка вывода")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	postalCodeFilter := r.URL.Query().Get("post")
	fmt.Println(postalCodeFilter)
	criticalityFilter := r.URL.Query().Get("crit")
	fmt.Println(criticalityFilter)

	if postalCodeFilter != "" || criticalityFilter != "" { //проверяем наличие хотябы одного чтобы стартануть, если квери параметров не было значит выгоняем все сразу
		if postalCodeFilter != "" {
			fmt.Println("Обнаружен queri параметр /почтовый индекс/ ")
			pCF, err := strconv.Atoi(postalCodeFilter)
			if err != nil {
				fmt.Println("Не удалось распарсить int")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			fmt.Println("Результат парсинга", pCF)
			fmt.Println("Значение в ")
			for _, value := range storageMessage {
				if pCF == value.PostalCode {
					fmt.Println("найдено совпадение", value)
					rM, err := json.Marshal(value)
					if err != nil {
						fmt.Println("Не получилось смаршалить все в JSON")
						w.WriteHeader(http.StatusInsufficientStorage)
						return
					}
					if _, err = w.Write(rM); err != nil {
						fmt.Println("Не получилось отправить в ответ", err)
						w.WriteHeader(http.StatusInsufficientStorage)
						return
					}

				} else {
					fmt.Println("Скипает")
					continue
				}
			}

		}
		if criticalityFilter != "" {
			fmt.Println("Обнаружен queri параметр /критичность письма/")
			cF, err := strconv.ParseBool(criticalityFilter)
			if err != nil {
				fmt.Println("Не удалось распарсить bool")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			for _, value := range storageMessage {
				if cF == value.Criticality {
					printJsonform(value, w)
				} else {
					continue
				}
			}
		}
	} else {
		fmt.Println("Фильтров не найдено")
		for _, value := range storageMessage {
			printJsonform(value, w)
		}
	}

}

func messageSearch(w http.ResponseWriter, r *http.Request) {
	httpRequest, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Не удалось обработать запрос", err)
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	numberId, err := strconv.Atoi(string(httpRequest))
	if err != nil {
		fmt.Println("Не удалось конвертировать в целочисленный тип", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	value, ok := storageMessage[numberId]
	if ok {
		fmt.Println("Обнаружено совпадение в поисковом запросе", value)
		rM, err := json.Marshal(value)
		if err != nil {
			fmt.Println("Не получилось обработать ответ на запрос")
			w.WriteHeader(http.StatusInsufficientStorage)
			return
		}
		if _, err = w.Write(rM); err != nil {
			fmt.Println("Не получилось отправить ответ на запрос", err)
			w.WriteHeader(http.StatusInsufficientStorage)
			return
		}

	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("Совпадений не найдено")
	}
}

func printJsonform(value Message, w http.ResponseWriter) {
	rM, err := json.Marshal(value)
	if err != nil {
		fmt.Println("Не получилось смаршалить все в JSON")
		w.WriteHeader(http.StatusInsufficientStorage)
		return
	}
	if _, err = w.Write(rM); err != nil {
		fmt.Println("Не получилось отправить в ответ", err)
		w.WriteHeader(http.StatusInsufficientStorage)
		return
	}
}
