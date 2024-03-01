package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Dice() int {
	rollvalue := rand.Intn(6) + 1
	return rollvalue
}

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	counts := make([]int, 13)
	const rolls = 10000
	var mute sync.Mutex

	wg.Add(rolls)
	for i := 0; i < rolls; i++ {
		go func() {
			defer wg.Done()
			dice1 := Dice()
			dice2 := Dice()
			fmt.Println("dice1 ", dice1)
			fmt.Println("dice2 ", dice2)

			sum := dice1 + dice2

			mute.Lock()
			counts[sum]++
			mute.Unlock()
		}()

	}
	wg.Wait()
	fmt.Println("Sum\tFrequency")
	for sum, count := range counts {
		if sum >= 2 {
			fmt.Printf("%d\t%d\n", sum, count)
		}
	}
}
