package models

import (
	"math/rand"
	"sync"
)

var dishNames = []string{"burger", "pizza", "popcorn", "cake", "ice-cream"}

type Dish struct {
	Name         string
	MorselsCount int
	sync.Mutex
}

func newDish(name string) *Dish {
	randomMorselCount := 6 + rand.Intn(5)
	dish := Dish{MorselsCount: randomMorselCount, Name: name}
	return &dish
}

func PrepareDishes() []Dish {
	var dishes []Dish
	for i := 0; i < 5; i++ {
		newDish := newDish(dishNames[i])
		dishes = append(dishes, *newDish)
	}
	return dishes
}
