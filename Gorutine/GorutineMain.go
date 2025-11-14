package Gorutine

import (
	"fmt"
	"sync"
	"time"
)

func GorutineMain() {

	var wg sync.WaitGroup

	wg.Add(3)

	go helloGorutine(1, &wg)
	go helloGorutine(2, &wg)
	go helloGorutine(3, &wg)

	wg.Wait()

	fmt.Println("\nDone")
}

func helloGorutine(number int, wg *sync.WaitGroup) {

	for i := 0; i < 5; i++ {

		fmt.Println("Я горутина", number, "делаю вывод на экран в", i+1, "раз")
		time.Sleep(1 * time.Second)

		if i == 4 {
			wg.Done()
		}
	}
}
