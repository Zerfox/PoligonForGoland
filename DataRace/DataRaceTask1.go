package DataRace

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
)

/*
Создать переменную (инт) - кол-во голосов за кандидата
Спровоцировать DataRace
Зафиксировать влияние гонки данных на результат
Решить гонку данных при помощи всего что я знаю
Зафиксировать отсутствие DataRace
*/

func DataRaceTask1Main() {

	var vote int64
	var atomicVote atomic.Int64
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	numberOfVotes := rand.Intn(1000)

	fmt.Println("Number of Votes: ", numberOfVotes, "Number of vote = ", vote)

	for i := 0; i < numberOfVotes; i = i + 2 {

		//go ballot(&vote, wg)

		go atomicBallot(&vote, wg)
		wg.Add(1)

		go mutexBallot(&vote, wg, mu)
		wg.Add(1)
	}
	for j := 0; j < numberOfVotes; j++ {
		go atomicVoteBallot(&atomicVote, wg)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println("Number of Votes: ", numberOfVotes, "\nNumber of vote = ", vote)
	fmt.Println("Number of атомарный Votes: ", numberOfVotes, "\nNumber of атомарный Vote = ", atomicVote.Load())
}

/*
При запуске функции ballot возникает серьезная утечка данных происходящая из-за DataRace.
К примеру данного когда при численности голосующих в 566 человек было  запущено в работу 566 горутин, что привело к потере 34 голосов, получив в результате 532 из 566.
Было проведено несколько тестов. Они показали, что датарэйс возникает независимо от количества вызванных горутин

Успешное выполнение:
 1. - atomicBallot выполняет тот же функционал, что и ballot,
    но полностью ликвидирует опасность возникновения dataRace на основе работы атомика, который ограничевает доступ к данным для других горутин до конечного выполнения функции,
    что позволяет гарантировать успешную обработку.
 2. - mutexBallot, блокирует доступ к данным для других горутин, до момента завершения, что приводит к формированию очереди для других горутин, и конечно полностью исключает DataRace
 3. - при перекрестной работе обоих методов все равно удается добиться желаемого результата без потерь.
 4. - для адаптивности и возможности работать с переменной при помощи не только атомиков,
    но еще и при помощи мьютекса я сделал переменную изначально 64 интом, но для наглядности добавил еще и обычную переменную атомика
*/
func mutexBallot(vote *int64, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	*vote++
	fmt.Println("Vote: ", *vote)
}

// работа с атомарной переменной
func atomicVoteBallot(vote *atomic.Int64, wg *sync.WaitGroup) {
	defer wg.Done()
	vote.Add(1)
	fmt.Println("Атомарный счетчик:", *vote)
}

// работа с обычной переменой инт 64
func atomicBallot(vote *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(vote, 1)
	fmt.Println("Счетчик:", atomic.LoadInt64(vote))
}

func ballot(vote *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	*vote++
	fmt.Println("ballot: ", *vote)
}
