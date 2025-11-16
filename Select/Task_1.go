package Select

import (
	"fmt"
	"strconv"
	"time"
)

func SelectMain() {

	intChan := make(chan int)
	stringChan := make(chan string)
	floatChan := make(chan float64)

	go func() {
		i := 0
		for {
			intChan <- i
			i++
			time.Sleep(100 * time.Millisecond)
		}

	}()

	go func() {
		i := 0
		for {
			stringChan <- "str_" + strconv.Itoa(i)
			i++
			time.Sleep(100 * time.Millisecond)
		}

	}()

	go func() {
		var float float64
		for {
			floatChan <- float
			float = float + 0.1
			time.Sleep(100 * time.Millisecond)
		}

	}()

	for {
		select {
		case iChan := <-intChan:
			fmt.Println(iChan)
		case strChan := <-stringChan:
			fmt.Println(strChan)
		case fChan := <-floatChan:
			fmt.Println(fChan)
		}
	}
}
