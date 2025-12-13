package DataRace

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
)

/*
В данном задании нужно было реализовать почтовик на который сыпяться сообщения и проверить будет ли гонка данных в этом случае и нужно было проверить есть ли она
Я прогнал это конкурентном варианте и у меня не появилось гонки данных, если я правильно понимаю, гонки не было из-за того что добавление данных в слайс не требует точности
Я пробовал начиная с 10 добавлять нули до 10000 и ни в одном случае поиск гонки данных -race не обнаружил у меня гонки
*/

func DataRaceTask2Main() {
	wg := &sync.WaitGroup{}
	mail := make([]string, 0, 100)
	ptMail := &mail

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		sendMessage(ptMail, wg)
	}

	wg.Wait()
	fmt.Println("Скрипт отработал: ", mail)
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

func sendMessage(mail *[]string, wg *sync.WaitGroup) {
	defer wg.Done()
	*mail = append(*mail, takeString())
}
