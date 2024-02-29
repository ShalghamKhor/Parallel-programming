package main

import (
	"fmt"
	"sync"
)

type Heap struct {
	heap []int
	lock sync.Mutex
	size int
}

func newHeap(cap int) *Heap {
	return &Heap{
		heap: make([]int, cap),
		size: 0,
	}
}

func (heap *Heap) insert(value int) {
	heap.lock.Lock()
	defer heap.lock.Unlock()
	if heap.size == len(heap.heap) {
		fmt.Println("Heap is full")
	}

	heap.heap[heap.size] = value
	current := heap.size
	heap.size++

	for current > 0 && heap.heap[current] < heap.heap[heap.getParent(current)] {
		heap.swap(current, heap.getParent(current))
		current = heap.getParent(current)
	}

}

func (heap *Heap) extractMin() (int, bool) {
	heap.lock.Lock()
	defer heap.lock.Unlock()

	if heap.size == 0 {
		return 0, false
	}

	minValue := heap.heap[0]
	heap.size--
	heap.heap[0] = heap.heap[heap.size]
	heap.down(0)
	return minValue, true
}

func (heap *Heap) swap(i, j int) {
	heap.heap[i], heap.heap[j] = heap.heap[j], heap.heap[i]
}

func (heap *Heap) getParent(i int) int {
	return (i - 1) / 2
}

func (heap *Heap) down(i int) {
	lastIndex := heap.size - 1

	for {
		leftChild := 2*i + 1
		rightChild := 2*i + 2
		smallest := i

		if leftChild <= lastIndex && heap.heap[leftChild] < heap.heap[smallest] {
			smallest = leftChild
		}

		if rightChild <= lastIndex && heap.heap[rightChild] < heap.heap[smallest] {
			smallest = rightChild
		}

		if smallest == i {
			break
		}

		heap.swap(i, smallest)
		i = smallest
	}
}

func (heap *Heap) getSize() int {
	heap.lock.Lock()
	defer heap.lock.Unlock()
	return heap.size
}

func main() {
	heap := newHeap(10)
	heap.insert(3)
	heap.insert(1)
	heap.insert(2)

	min, ok := heap.extractMin()
	fmt.Println("size of the heap: ", heap.getSize())

	if ok {
		fmt.Println("Min vakue : ", min)
	} else {
		fmt.Println("Heap is empty: ")
	}

	fmt.Println("Size of the min heap: ", heap.getSize())
}
