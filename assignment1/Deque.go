package assignment1

import (
	"fmt"
	"sync"
)

type Node struct {
	Nodes int
	prev  *Node
	next  *Node
}

type Deque struct {
	head *Node
	tail *Node
	lock sync.Mutex
}

func (dq *Deque) isEmpty() bool {
	dq.lock.Lock()
	defer dq.lock.Unlock()
	return dq.head == nil
}

func (dq *Deque) addFirst(node int) {
	dq.lock.Lock()
	defer dq.lock.Unlock()
	newNode := &Node{
		Nodes: node,
		prev:  nil,
		next:  nil,
	}
	if dq.head == nil {
		dq.head = newNode
		dq.tail = newNode
	} else {
		newNode.next = dq.head
		dq.head.prev = newNode
		dq.head = newNode
	}

}

func (dq *Deque) addLast(node int) {
	dq.lock.Lock()
	defer dq.lock.Unlock()
	newNode := &Node{
		Nodes: node,
		prev:  nil,
		next:  nil,
	}

	if dq.tail == nil {
		dq.head = newNode
		dq.tail = newNode
	} else {
		dq.tail.next = newNode
		newNode.prev = dq.tail
		dq.tail = newNode
	}

}

func (dq *Deque) removeFirst() int {
	dq.lock.Lock()
	defer dq.lock.Unlock()
	if dq.head == nil {
		return 0
	}

	node := dq.head.Nodes
	if dq.head == dq.tail {
		dq.head = nil
		dq.tail = nil
	} else {
		dq.head = dq.head.next
		dq.head.prev = nil
	}

	return node
}

func (dq *Deque) removeLast() int {
	dq.lock.Lock()
	defer dq.lock.Unlock()
	if dq.tail == nil {
		return 0
	}

	node := dq.tail.Nodes
	if dq.tail == dq.head {
		dq.head = nil
		dq.tail = nil
	} else {
		dq.tail = dq.tail.prev
		dq.tail.next = nil
	}

	return node
}

func (dq *Deque) Display() {
	current := dq.head
	for current != nil {
		fmt.Print(current.Nodes, " ")
		current = current.next
	}
	fmt.Println()
}

func Task2() {
	dq := &Deque{}
	dq.addLast(1)
	dq.addLast(2)
	dq.addFirst(0)
	dq.Display() // Expected: 0 1 2

	// Deleting elements
	value := dq.removeFirst()
	fmt.Println("Deleted from front:", value) // Expected: 0

	value1 := dq.removeLast()
	fmt.Println("Deleted from back:", value1) // Expected: 2
	b := dq.isEmpty()
	fmt.Println(b)
	dq.Display()
}
