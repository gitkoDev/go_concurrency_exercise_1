package models

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var friendsNames = []string{"Denis", "Kate", "Cocos", "Baby"}
var Dishes = PrepareDishes()

type Friend struct {
	name string
}

func GatherFriends() []Friend {
	var friends []Friend
	for i := 0; i < 4; i++ {
		friends = append(friends, Friend{name: friendsNames[i]})
	}

	return friends
}

func (f *Friend) EatMorsel(wg *sync.WaitGroup, foodChan chan int) {

	// Repeat the cycle of picking a random dish to eat until all dishes are eaten, then return
	for {
		randTime := rand.Intn(5) + 1
		randomNum := rand.Intn(len(Dishes))

		// 1. Pick a dish to eat and lock its mutex
		Dishes[randomNum].Lock()

		// 2. If the dish is already empty, we unlock the mutex and repeat the loop to pick another one
		if Dishes[randomNum].MorselsCount == 0 {
			Dishes[randomNum].Unlock()
			continue
		}

		// 3. If there's still morsels left, eat one
		fmt.Printf("%s is eating a morsel of %s. %d morsels left\n", f.name, Dishes[randomNum].Name, Dishes[randomNum].MorselsCount-1)
		Dishes[randomNum].MorselsCount--

		// 4. If the dish is empty after we ate a morsel, send a signal into the channel
		if Dishes[randomNum].MorselsCount == 0 {
			fmt.Printf("----------%s is finished----------\n", Dishes[randomNum].Name)
			wg.Done()

			foodChan <- 1

			Dishes[randomNum].Unlock()
			time.Sleep(time.Second * time.Duration(randTime))
			continue

		}
		time.Sleep(time.Second * time.Duration(randTime))
		Dishes[randomNum].Unlock()

	}

}
