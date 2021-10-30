package stack

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Element generic.Type

type Stack struct {
	Elements []Element
	lock     sync.RWMutex
}

func New() *Stack {
	s := &Stack{}
	s.Elements = []Element{}

	return s
}

func (s *Stack) Push(element Element) {
	s.lock.Lock()
	s.Elements = append(s.Elements, element)
	s.lock.Unlock()
}

func (s *Stack) Pop() *Element {
	s.lock.Lock()
	poppedElement := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[0 : len(s.Elements)-1]
	s.lock.Unlock()
	return &poppedElement
}

func (s *Stack) Size() int {
	return len(s.Elements)
}

func (s *Stack) IsEmpty() bool {
	return len(s.Elements) <= 0
}
