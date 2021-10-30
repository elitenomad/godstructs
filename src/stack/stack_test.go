package stack

import "testing"

func TestPush(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	if size := s.Size(); size != 4 {
		t.Errorf("Expected count of 4 and got %d ", size)
	}
}

func TestPop(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(1)
	s.Push(1)
	s.Pop()
	if size := len(s.Elements); size != 2 {
		t.Errorf("Expected count of 2 and got %d ", size)
	}

	s.Pop()
	s.Pop()
	if size := len(s.Elements); size != 0 {
		t.Errorf("Expected count of 0 and got %d", size)
	}

	if !s.IsEmpty() {
		t.Errorf("IsEmpty should return true")
	}
}

func TestSize(t *testing.T) {
	s := New()
	if size := len(s.Elements); size != 0 {
		t.Errorf("Expected count of 0 and got %d", size)
	}

	s.Push(1)
	s.Push(2)
	if size := len(s.Elements); size != 2 {
		t.Errorf("Expected count of 2 and got %d", size)
	}
}
