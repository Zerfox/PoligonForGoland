package Gorutine

import (
	"fmt"
	"math/rand"
)

func Task_2() {
	transferPoint := make(chan int, 10)
	var outInt int

	go takeInt(transferPoint)
	go takeInt(transferPoint)
	go takeInt(transferPoint)
	go takeInt(transferPoint)
	go takeInt(transferPoint)
	go takeInt(transferPoint)
	go takeInt(transferPoint)
	go takeInt(transferPoint)

	for i := 0; i <= 7; i++ {
		outInt += <-transferPoint
		fmt.Println(outInt)
	}

}

func takeInt(point chan int) {
	number := rand.Intn(100)
	point <- number

}
