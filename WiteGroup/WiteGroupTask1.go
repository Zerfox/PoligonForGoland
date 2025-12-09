package WiteGroup

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/* Задача: реализовать работу садоводов, которые следят за большим огородом
Особенности: 1) время полива случайное от 500 мс до 1000 мс
			 2) кол-во огородников тоже случайное
			 3) завязать работу горутин на waitGroup
			 4) Вывести в конце сообщение об окончании выполнения программы

*/

func WiteGroupTask1() {
	wg := sync.WaitGroup{}
	gardians := rand.Intn(10)

	for i := 0; i < gardians; i++ {
		wg.Add(1)
		go pour(i+1, &wg)
	}

	wg.Wait()
}

func pour(numberGarian int, wg *sync.WaitGroup) {
	defer wg.Done()

	timeout := 500 + (100 * (rand.Intn(5)))

	fmt.Println("Начало полива. Садовод № ", numberGarian)
	time.Sleep(time.Duration(timeout) * time.Millisecond)
	fmt.Println("Полив окончен. Садовод № ", numberGarian)

}
