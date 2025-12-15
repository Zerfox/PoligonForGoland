package DataRace

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
)

/*
В данном задании нужно было реализовать почтовик на который сыпяться сообщения и проверить будет ли гонка данных в этом случае и нужно было проверить есть ли она
	Гонка данных действительно возникает, я пофиксил ее мьютексом, пушто атомики требуют больших затрат
*/

func DataRaceTask2Main() {
	wg := &sync.WaitGroup{}
	mail := make([]string, 0, 100)
	ptMail := &mail
	mtx := &sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go sendMessage(ptMail, wg, mtx)
	}

	wg.Wait()
	fmt.Println("Скрипт отработал: ", len(mail))
}

func takeString() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "
	sb := strings.Builder{}
	sb.Grow(5)
	for i := 0; i < 10; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}

	return sb.String()
}

func sendMessage(mail *[]string, wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()
	mtx.Lock()
	*mail = append(*mail, takeString())
	mtx.Unlock()
}
