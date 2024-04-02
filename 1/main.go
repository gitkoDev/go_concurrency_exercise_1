package main

import (
	"fmt"
	"sync"

	"github.com/gitkoDev/go-testing/modelsWg"
	models "github.com/gitkoDev/go-testing/modelschan"
)

func main() {
	// doChannels()
	doWaitGroups()
}

// Channels approach
func doChannels() {
	people := []models.Person{
		{Name: "Alice", ReadinessChan: make(chan bool)},
		{Name: "Bob", ReadinessChan: make(chan bool)},
	}

	var aliceReadyChan = people[0].ReadinessChan
	var bobReadyChan = people[1].ReadinessChan

	fmt.Println("Let's go for a walk")

	// Get dressed
	for _, person := range people {
		go person.GetReady()
	}

	select {
	case <-aliceReadyChan:
		<-bobReadyChan
	case <-bobReadyChan:
		<-aliceReadyChan
	}

	// Set the alarm and go
	for _, person := range people {
		go person.PutShoesOn()
	}
}

// Waitgroups approach
func doWaitGroups() {
	prepare()
	getShoesOn()
}

func prepare() {
	var wg sync.WaitGroup

	people := []modelsWg.Person{
		{Name: "Alice"},
		{Name: "Bob"},
	}

	for i := 0; i < 2; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			people[i].GetReady()
		}()
	}

	wg.Wait()
}

func getShoesOn() {
	var wg sync.WaitGroup

	people := []modelsWg.Person{
		{Name: "Alice"},
		{Name: "Bob"},
	}

	for i := 0; i < 2; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			people[i].PutShoesOn()
		}()
	}

	wg.Wait()
}
