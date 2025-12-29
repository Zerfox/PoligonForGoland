package Cuncurenci

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func CimulatorMetCentraMain() {

	wg := &sync.WaitGroup{}
	// первичная инициализация верховного контекста parentCancel
	parentContext, _ := context.WithCancel(context.Background())

	// три контекста для каждого вида датчиков
	tempirSCtx, tempirSCancel := context.WithCancel(parentContext)     // контекст для контроля датчиков температуры
	seismicSCtx, seismicSCancel := context.WithCancel(parentContext)   // контекст для контроля датчиков сейсмической активности
	humiditySCtx, humiditySCancel := context.WithCancel(parentContext) //контекст для контроля датчиков влажности воздуха

	for {
		////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		wg.Add(1)
		go takeSensTemp(tempirSCtx, wg)
		wg.Add(1)
		go takeSensSeysmetic(seismicSCtx, wg)
		wg.Add(1)
		go takeSensVlasjnost(humiditySCtx, wg)

		fmt.Println("Введите значение 1 отключит темпиратурный датчик, 2 - сейсмичский, 3 - влажности воздуха")

		scanner := bufio.NewScanner(os.Stdin) // инициализация сканера
		if ok := scanner.Scan(); !ok {        // быстрая проверка на наличие качественного скана "ок", если не ок, то выдаст ошибку
			fmt.Println("Ошибка ввода")
			return
		}
		input := scanner.Text()
		userRequest, err := strconv.Atoi(input) // конвертация с строки в целочисленный тип данных
		if err != nil {
			fmt.Println("Ошибка ввода данных")
		}

		switch userRequest {
		case 1:
			fmt.Println("Вывожу из работы датчик температуры")
			tempirSCancel()
		case 2:
			fmt.Println("Вывожу из работы  датчик сейсмической активности ")
			seismicSCancel()
		case 3:
			fmt.Println("Вывожу из работы датчик влажности воздуха ")
			humiditySCancel()
		case 4:
			fmt.Println("Продолжаю работу")

		}
	}

	wg.Wait()
}
