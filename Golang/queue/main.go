package main

import "fmt"

type Queue struct {
	num []int
	maxCapacity int // Set a maximum capacity for the queue
}

func (e *Queue) Enqueue(val int) {
	if !e.isFull() {
		e.num = append(e.num, val)
	} else {
		fmt.Println("Queue is full; cannot enqueue.")
	}
}

func (e *Queue) Dequeue() int {
	if len(e.num) == 0 {
		return -1 // Queue is empty
	}

	dequeue := e.num[0]
	e.num = e.num[1:]
	return dequeue
}

func (e *Queue) isFull() bool {
	return len(e.num) >= e.maxCapacity
}

func main() {
	// Initialize the queue with a max capacity
	enqueue := Queue{maxCapacity: 3}

	enqueue.Enqueue(1)
	enqueue.Enqueue(2)
	enqueue.Enqueue(4)

	enqueue.Enqueue(12)

	// Test dequeue
	enqueue.Dequeue()
	enqueue.Dequeue()

	fmt.Println(enqueue) 
}

