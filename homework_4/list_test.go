package homework_4

import "testing"

func TestNewListSize(t *testing.T) {
	list := NewList()

	expected := 0

	if result := list.Size(); result != expected {
		t.Fatalf("bad size for new List: got %d, expected %d", result, expected)
	}
}

func TestNewListEmpty(t *testing.T) {
	list := NewList()

	if result := list.Empty(); result != true {
		t.Fatalf("bad empty for new List: got %v, expected %v", result, true)
	}
}

func TestNewListNodeNil(t *testing.T) {
	list := NewList()

	if result := list.Front(); nil != result {
		t.Fatalf("bad front for new List: got %v, expected %v", result, nil)
	}
}
