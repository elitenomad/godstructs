package set

import "github.com/cheekybits/genny/generic"

type Element generic.Type

type Set struct {
	Elements map[Element]bool
}

func (s *Set) Add(element Element) *Set {
	if s.Elements == nil {
		s.Elements = make(map[Element]bool)
	}

	if _, found := s.Elements[element]; !found {
		s.Elements[element] = true
	}

	return s
}

func (s *Set) Clear() {
	s.Elements = make(map[Element]bool)
}

func (s *Set) Delete(element Element) bool {
	if _, found := s.Elements[element]; found {
		delete(s.Elements, element)
		return true
	}

	return false
}

func (s *Set) Has(element Element) bool {
	_, found := s.Elements[element]
	return found
}

func (s *Set) List() []Element {
	els := []Element{}
	for i := range s.Elements {
		els = append(els, s.Elements[i])
	}

	return els
}

func (s *Set) Size() int {
	return len(s.Elements)
}
