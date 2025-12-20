package RWmutex

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

/*
В данном задании предомной стояла задача наглядно изучить разницу между обычным мъютексом и его
RWMuteх братом. В Результате я понял, что РВ мъютекс хорош только в тех случаях, когда операций с чтением значительно больше, нежели с запиьсю,
но когда намного больше записей, используем обычный мъютекс.
*/
func RWmutexMine() {

	storageKey := make([]string, 0, 30)
	storage := make(map[string]string)

	wg := &sync.WaitGroup{}
	defer wg.Wait()

	start := time.Now()

	mtx := &sync.RWMutex{}
	//rwmtx := sync.RWMutex{}

	for i := 0; i < 3000; i++ {
		wg.Add(1)
		go writer(storage, mtx, wg, &storageKey)

	}
	//time.Sleep(2 * time.Second)

	for i := 0; i < 300000; i++ {
		wg.Add(1)
		go reader(storage, mtx, wg, storageKey)
	}

	wg.Wait()

	duration := time.Since(start)
	fmt.Println(duration.Seconds())

}

func writer(storage map[string]string, mtx *sync.RWMutex, wg *sync.WaitGroup, storageKey *[]string) {
	defer wg.Done()

	key := takeString()
	value := takeString()

	mtx.Lock()
	storage[key] = value
	*storageKey = append(*storageKey, key)
	mtx.Unlock()
}

func reader(storage map[string]string, mtx *sync.RWMutex, wg *sync.WaitGroup, storageKey []string) {

	defer wg.Done()
	randomKey := storageKey[rand.Intn(len(storageKey))]

	//fmt.Println(randomKey)

	mtx.RLock()
	var _ string = storage[randomKey]
	defer mtx.RUnlock()
}

func takeString() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "
	sb := strings.Builder{}
	sb.Grow(5)
	for i := 0; i < 3; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}

	return sb.String()
}
