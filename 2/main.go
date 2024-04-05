package main

import (
	"fmt"
	"sync"
	"test/Desktop/Coding/go_projects/concurrency/2/models"
)

func main() {
	// Initialize dishes and friends
	friends := models.GatherFriends()
	foodChan := make(chan int)
	dishesEaten := 0

	for _, dish := range models.Dishes {
		fmt.Printf("---%s: %d morsels\n", dish.Name, dish.MorselsCount)
	}
	fmt.Println("---------------------------")

	// Init wait group
	var wg sync.WaitGroup

	wg.Add(5)

	for index, _ := range friends {
		go friends[index].EatMorsel(&wg, foodChan)
	}

	for {
		result := <-foodChan
		dishesEaten += result

		if dishesEaten == len(models.Dishes) {
			break
		}
	}

	wg.Wait()
	fmt.Println("finished")
}
