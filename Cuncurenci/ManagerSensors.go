package Cuncurenci

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Так на часах 2:40, а мне еще сегодня идти на работу, пишу памятку.
Система работает, на этапе получения данных от датчиков и завершается штатно, теперь нужно:
	1) Подключить систему контекстного отключения
	2)
	3)
*/

func collectionBySensor(sensorConst float32, wg *sync.WaitGroup) float32 {
	defer wg.Done()
	storageNum := make([]float32, 0, 5)
	sens := rand.Intn(10000)

	for i := 0; i <= sens; i++ {
		if sens == 0 {
			fmt.Println("не удалось получить данные с сенсоров")
			break
		}
		num := sensorConst * float32(rand.Intn(1000))
		storageNum = append(storageNum, num)
	}

	var averageValue float32 = 0.0

	for i := 0; i < len(storageNum); i++ {
		averageValue = averageValue + storageNum[i]
		if i+1 == len(storageNum) {
			averageValue = averageValue / float32(len(storageNum))
		}
	}
	fmt.Println("\nСреднее значение сенсоров", averageValue, " Кол-во сенсоров: ", len(storageNum))

	return averageValue

}

func takeSensTemp(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	tPointT := make(chan float32)
	var tempirS float32 = 0.75

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Датчик температуры отключен")
			return
		default:
			wg.Add(1)
			temper := collectionBySensor(tempirS, wg)
			fmt.Println("Температура воздуха", temper)
			tPointT <- temper
			time.Sleep(3 * time.Second)
		}
	}

}

func takeSensSeysmetic(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	sPointS := make(chan float32)
	var seismicS float32 = 30.2

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Датчик сейсмической активности отключен")
			return
		default:
			wg.Add(1)
			number := collectionBySensor(seismicS, wg)
			fmt.Println("Уровень сейсмической активности составляет ", number)
			sPointS <- number
			time.Sleep(3 * time.Second)
		}
	}

}

func takeSensVlasjnost(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	hPointH := make(chan float32)
	var humidityS float32 = 1.25

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Датчик влажности воздуха  отключен")
			return
		default:
			number := collectionBySensor(humidityS, wg)
			fmt.Println("Влажности воздуха", number)
			hPointH <- number
			time.Sleep(3 * time.Second)
		}
	}

}
