package set

import (
	"fmt"
	"testing"
)

func populateSet(count int, start int) *Set {
	set := Set{}
	for i := start; i < (start + count); i++ {
		set.Add(fmt.Sprintf("element%d", i))
	}
	return &set
}

func TestAdd(t *testing.T) {
	set := populateSet(3, 0)
	if size := set.Size(); size != 3 {

		t.Errorf("Expected value of 3 and got %d", size)
	}
	set.Add("element1") //should not add it, already there
	if size := set.Size(); size != 3 {
		t.Errorf("Expected value of 3 and got %d", size)
	}
	set.Add("element4") //should not add it, already there
	if size := set.Size(); size != 4 {
		t.Errorf("Expected value of 4 and got %d", size)
	}
}

func TestClear(t *testing.T) {
	set := populateSet(3, 0)
	set.Clear()
	if size := set.Size(); size != 0 {
		t.Errorf("Expected value of 0 and got %d", size)
	}
}

func TestDelete(t *testing.T) {
	set := populateSet(3, 0)
	set.Delete("element2")
	if size := set.Size(); size != 2 {
		t.Errorf("Expected value of 2 and got %d", size)
	}
}

func TestHas(t *testing.T) {
	set := populateSet(3, 0)
	has := set.Has("element2")
	if !has {
		t.Errorf("expected element2 to be there")
	}
	set.Delete("element2")
	has = set.Has("element2")
	if has {
		t.Errorf("expected element2 to be removed")
	}
	set.Delete("element1")
	has = set.Has("element1")
	if has {
		t.Errorf("expected element1 to be removed")
	}
}

func TestItems(t *testing.T) {
	set := populateSet(3, 0)
	elements := set.List()
	if len(elements) != 3 {
		t.Errorf("wrong count, expected 3 and got %d", len(elements))
	}
	set = populateSet(520, 0)
	elements = set.List()
	if len(elements) != 520 {
		t.Errorf("wrong count, expected 520 and got %d", len(elements))
	}
}

func TestSize(t *testing.T) {
	set := populateSet(3, 0)
	elements := set.List()
	if len(elements) != set.Size() {
		t.Errorf("wrong count, expected %d and got %d", set.Size(), len(elements))
	}
	set = populateSet(0, 0)
	elements = set.List()
	if len(elements) != set.Size() {
		t.Errorf("wrong count, expected %d and got %d", set.Size(), len(elements))
	}
	set = populateSet(10000, 0)
	elements = set.List()
	if len(elements) != set.Size() {
		t.Errorf("wrong count, expected %d and got %d", set.Size(), len(elements))
	}
}
