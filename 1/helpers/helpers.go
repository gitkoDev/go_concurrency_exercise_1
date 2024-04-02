package helpers

import "math/rand"

// Helpers
func GetRandTime() int {
	// Generate number 10-15
	randTime := rand.Intn(6)
	// return 10 + randTime
	return randTime
}
