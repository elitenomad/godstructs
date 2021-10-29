package queue

import "testing"

func TestEnqueue(t *testing.T) {
	s := New()
	s.Enqueue(1)
	s.Enqueue(2)
	s.Enqueue(3)
	s.Enqueue(4)

	if size := s.Size(); size != 4 {
		t.Errorf("Expected count of 4 and got %d ", size)
	}
}

func TestDequeue(t *testing.T) {
	s := New()
	s.Enqueue(1)
	s.Enqueue(1)
	s.Enqueue(1)
	s.Dequeue()
	if size := len(s.Elements); size != 2 {
		t.Errorf("Expected count of 2 and got %d ", size)
	}

	s.Dequeue()
	s.Dequeue()
	if size := len(s.Elements); size != 0 {
		t.Errorf("Expected count of 0 and got %d", size)
	}

	if !s.IsEmpty() {
		t.Errorf("IsEmpty should return true")
	}
}

func TestIsEmpty(t *testing.T) {
	s := New()
	if !s.IsEmpty() {
		t.Errorf("IsEmpty is expected to return true")
	}

	s.Enqueue(1)
	if s.IsEmpty() {
		t.Errorf("IsEmpty is expected to return false")
	}

	s.Dequeue()
	if !s.IsEmpty() {
		t.Errorf("IsEmpty is expected to return true")
	}
}

func TestSize(t *testing.T) {
	s := New()
	if size := len(s.Elements); size != 0 {
		t.Errorf("Expected count of 0 and got %d", size)
	}

	s.Enqueue(1)
	s.Enqueue(2)
	if size := len(s.Elements); size != 2 {
		t.Errorf("Expected count of 2 and got %d", size)
	}
}

func TestFront(t *testing.T) {
	s := New()

	s.Enqueue(1)
	s.Enqueue(2)
	if element := s.Front(); *element != 1 {
		t.Errorf("Expected value of 1  and got %d", element)
	}
}

func TestRear(t *testing.T) {
	s := New()

	s.Enqueue(1)
	s.Enqueue(2)
	if element := s.Rear(); *element != 2 {
		t.Errorf("Expected value of 2 and got %d", element)
	}
}
