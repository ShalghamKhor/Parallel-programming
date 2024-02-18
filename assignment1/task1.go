package assignment1

import (
	"math/rand"
	"time"
)

func Dice(n int) {

	counts := make([]int, 13)

	rand.Seed(time.Now().UnixNano())

	const rolls = 1000

	for i := 0; i < rolls; i++ {
		dic1 := rand.Intn(6) + 1

	}
}
