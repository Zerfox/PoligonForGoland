package DataRace

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
)

func DataRaceTask2Main() {
	wg := &sync.WaitGroup{}
	mail := make([]string, 100)

	send(mail, wg)

	wg.Wait()
	fmt.Println(mail)
}

func takeOpinion() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "
	sb := strings.Builder{}
	sb.Grow(5)
	for i := 0; i < 10; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func send(mail []string, wg *sync.WaitGroup) {
	defer wg.Done()

	mail = append(mail, takeOpinion())
}
