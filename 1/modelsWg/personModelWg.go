package modelsWg

import (
	"fmt"
	"time"

	"github.com/gitkoDev/go-testing/helpers"
)

type Person struct {
	Name string
}

// Step 1
func (p *Person) GetReady() {
	// Set timer
	start := time.Now()
	// Do stuff
	p.closeWindows()
	p.getBelt()
	p.getGlasses()
	p.getKeys()
	p.getPhone()
	p.turnOffFans()
	// Check timer for the person
	timeElapsed := time.Since(start).Seconds()
	fmt.Printf("%s is ready!. It took %s %v seconds to get ready\n", p.Name, p.Name, int(timeElapsed))
}

// Step 2
func (p *Person) PutShoesOn() {
	// Set timer
	fmt.Println("------------")
	fmt.Println("Starting the alarm")
	fmt.Printf("%s is going to put the shoes om\n", p.Name)
	start := time.Now()
	fmt.Println("------------")
	// Do stuff
	time.Sleep(time.Second * time.Duration(helpers.GetRandTime()))
	// Check timer for the person
	timeElapsed := time.Since(start).Seconds()
	fmt.Printf("%s has put the shoes on!. It took %s %v seconds to do it\n", p.Name, p.Name, int(timeElapsed))
}

// Stuff to do
func (p *Person) getGlasses() {
	time.Sleep(time.Second * time.Duration(helpers.GetRandTime()))
	fmt.Printf("%s got glasses\n", p.Name)
}

func (p *Person) getBelt() {
	time.Sleep(time.Second * time.Duration(helpers.GetRandTime()))
	fmt.Printf("%s got belt\n", p.Name)
}

func (p *Person) closeWindows() {
	time.Sleep(time.Second * time.Duration(helpers.GetRandTime()))
	fmt.Printf("%s closed windows\n", p.Name)
}

func (p *Person) turnOffFans() {
	time.Sleep(time.Second * time.Duration(helpers.GetRandTime()))
	fmt.Printf("%s turned fans off\n", p.Name)
}

func (p *Person) getPhone() {
	time.Sleep(time.Second * time.Duration(helpers.GetRandTime()))
	fmt.Printf("%s got phone\n", p.Name)
}

func (p *Person) getKeys() {
	time.Sleep(time.Second * time.Duration(helpers.GetRandTime()))
	fmt.Printf("%s got keys\n", p.Name)
}
