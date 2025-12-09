package DataRace

import (
	"fmt"
	"math/rand"
	"sync"
)

func DataRaceTask1Main() {

	vote := 0
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	numberOfVotes := rand.Intn(100)

	fmt.Println("Number of Votes: ", numberOfVotes, "Number of vote = ", vote)

	for i := 0; i < numberOfVotes; i++ {
		wg.Add(1)
		go ballot(&vote)
	}

}
func ballot(vote *int) {
	*vote++
}
