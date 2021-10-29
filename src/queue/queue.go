package queue

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Element generic.Type

type Queue struct {
	Elements []Element
	lock     sync.RWMutex
}

func New() *Queue {
	q := &Queue{}
	q.Elements = []Element{}
	return q
}

func (q *Queue) Enqueue(element Element) {
	q.lock.Lock()
	q.Elements = append(q.Elements, element)
	q.lock.Unlock()
}

func (q *Queue) Dequeue() *Element { // O(N) (Look for Linked List implementation)
	q.lock.Lock()
	removedElement := q.Elements[0]
	q.Elements = q.Elements[1:]
	q.lock.Unlock()
	return &removedElement
}

func (q *Queue) Front() *Element {
	q.lock.RLock()
	element := q.Elements[0]
	q.lock.RUnlock()
	return &element
}

func (q *Queue) Rear() *Element {
	q.lock.RLock()
	element := q.Elements[len(q.Elements)-1]
	q.lock.RUnlock()
	return &element
}

func (q *Queue) Size() int {
	return len(q.Elements)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) <= 0
}
