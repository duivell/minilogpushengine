package queue

import (
	"fmt"
	"sync"
)

type Queue struct {
	mu    sync.Mutex
	items []int
	cap   int
}

func New(cap int) *Queue {
	return &Queue{cap: cap}
}

func (q *Queue) Enqueue(value int) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) >= q.cap {
		return fmt.Errorf("queue full")
	}

	q.items = append(q.items, value)
	return nil
}

func (q *Queue) Dequeue() (int, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) == 0 {
		return 0, fmt.Errorf("queue empty")
	}

	value := q.items[0]
	q.items = q.items[1:]
	return value, nil
}

func (q *Queue) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items) == 0
}
