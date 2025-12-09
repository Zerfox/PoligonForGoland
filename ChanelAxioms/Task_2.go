package ChanelAxioms

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func ChanelMain() {
	transferPoint := make(chan string)
	board := make([]string, 0, 10)
	n := takePeople() // кол-во людей от 0 до 10

	go interviewer(transferPoint, n)

	for i := 0; i < n; i++ {
		v, ok := <-transferPoint
		if !ok {
			fmt.Println("Весь список людей пройден")
			break
		}
		board = append(board, v)
		fmt.Println("Мнение №", i+1, ":", v)
	}

	fmt.Println("Кол-во мнений", n, "\nСами мнения:", board)

}

func interviewer(point chan string, n int) {

	pollingTimer := []int{300, 400, 500, 600, 700}

	for i := 0; i < n; i++ {
		point <- takeOpinion()
		duration := rand.Intn(5)
		time.Sleep(time.Duration(pollingTimer[duration]) * time.Millisecond)
	}

}

func takePeople() int {
	return rand.Intn(10)
}

func takeOpinion() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	sb := strings.Builder{}
	sb.Grow(5)
	for i := 0; i < 5; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}
