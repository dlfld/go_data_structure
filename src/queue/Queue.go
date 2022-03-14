package queue

import "sync"

type item interface {
}

// ItemQueue Item the type of the queue
type ItemQueue struct {
	items []item
	sync.RWMutex
}

type itemQueuer interface {
	NewQueue() ItemQueue
	Offer(t item)
	Poll() item
	IsEmpty() bool
	Size() int
	Peek() item
}

// NewQueue creates a new ItemQueue
func NewQueue() *ItemQueue {
	s := &ItemQueue{items: []item{}}
	return s
}

// Offer adds an Item to the end of the queue
func (s *ItemQueue) Offer(t item) {
	s.Lock()
	s.items = append(s.items, t)
	defer s.Unlock()
}

// Poll dequeue
func (s *ItemQueue) Poll() item {
	s.Lock()
	if s.IsEmpty() {
		panic("queue is empty")
	}
	t := s.items[0] // 先进先出
	if len(s.items) <= 1 {
		s.items = []item{}
		return t
	}
	s.items = s.items[1:len(s.items)]
	defer s.Unlock()
	return t
}

// IsEmpty Determine whether the queue is empty
func (s *ItemQueue) IsEmpty() bool {
	return s == nil || len(s.items) == 0
}

// Size returns the number of Items in the queue
func (s *ItemQueue) Size() int {
	return len(s.items)
}

// Peek return top of queue not remove
func (s *ItemQueue) Peek() (t item) {
	s.Lock()
	if s.IsEmpty() {
		panic("queue is empty")
	}
	t = s.items[0]
	defer s.Unlock()
	return t
}
