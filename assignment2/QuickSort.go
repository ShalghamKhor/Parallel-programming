package assignment2

import (
	"fmt"
	"sync"
)

// partition function that rearranges the array around a pivot
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // Swap elements
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1] // Move pivot to the correct position
	return i + 1
}

// worker function that performs the sorting
func quicksortWorker(arr []int, tasks chan [2]int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		low, high := task[0], task[1]
		fmt.Printf("Worker started task: sorting from index %d to %d\n", low, high)
		if low < high {
			pi := partition(arr, low, high)

			// Add new tasks for the left and right sub-arrays
			wg.Add(2) // Account for the new tasks
			tasks <- [2]int{low, pi - 1}
			tasks <- [2]int{pi + 1, high}
		}
		fmt.Printf("Worker completed task: sorting from index %d to %d\n", low, high)
	}
	fmt.Println("Worker exiting")
}

// main concurrent quicksort function
func concurrentQuicksort(arr []int) {
	tasks := make(chan [2]int, len(arr))
	var wg sync.WaitGroup

	// Start worker goroutines
	numWorkers := 4
	for i := 0; i < numWorkers; i++ {
		wg.Add(1) // Add 1 for each worker
		go quicksortWorker(arr, tasks, &wg)
	}

	// Initial task to sort the entire array
	wg.Add(1)
	fmt.Println("Main: adding initial task to sort the entire array")
	tasks <- [2]int{0, len(arr) - 1}

	// Close the tasks channel after all tasks have been sent
	go func() {
		wg.Wait() // Wait for all workers to finish their work
		fmt.Println("Main: all tasks completed, closing channel")
		close(tasks)
	}()

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Main: all goroutines finished")
}

func QuickSort() {
	arr := []int{33, 10, 55, 19, 27, 44, 99, 11}
	fmt.Println("Unsorted array:", arr)

	// Perform the concurrent QuickSort
	concurrentQuicksort(arr)

	fmt.Println("Sorted array:", arr)
}
