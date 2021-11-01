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

func TestListPushBackBasic(t *testing.T) {
	list := NewList()
	node := Node{12, nil, nil}

	resultNode := list.PushBack(&node)

	if resultNode.Value != node.Value {
		t.Fatalf("bad push back: got %v, expected %v", resultNode.Value, node.Value)
	}

	expectedSize := 1

	if expectedSize != list.Size() {
		t.Fatalf("bad push back: got %v, expected %v", list.Size(), expectedSize)
	}
}

func TestListBackEmpty(t *testing.T) {
	list := NewList()
	resultNode := list.Back()

	if resultNode != nil {
		t.Fatalf("bad back: got %v, expected %v", resultNode, nil)
	}
}

func TestListPushBackCompareBack(t *testing.T) {
	list := NewList()
	node := Node{12, nil, nil}

	resultNode := list.PushBack(&node)
	expectedNode := list.Back()

	if resultNode != expectedNode {
		t.Fatalf("bad push back and back: got %v, expected %v", resultNode, expectedNode)
	}
}

func TestListSomePushBack(t *testing.T) {
	list := NewList()

	for i := 0; i < 3; i++ {
		list.PushBack(&Node{i, nil, nil})
	}

	nodeBack := list.Back()

	if nodeBack == nil {
		t.Fatalf("bad push back: got %v", nodeBack)
	}

	node := list.Node
	for expected := 0; expected < 3; expected++ {
		if node.Value != expected {
			t.Fatalf("bad push back: got %v, expected %v", node.Value, expected)
		}
		node = node.Next
	}
}

func TestListRemoveHead(t *testing.T) {
	list := NewList()

	node := Node{1, nil, nil}
	list.PushBack(&node)

	expected := 1
	if list.Size() != expected {
		t.Fatalf("bad push back: got %v, expected %d", list.Size(), expected)
	}

	list.Remove(&node)
	expected = 0
	if list.Size() != expected {
		t.Fatalf("bad remove: got %v, expected %d", list.Size(), expected)
	}
}

func TestListRemoveTail(t *testing.T) {
	list := NewList()

	node1 := Node{1, nil, nil}
	list.PushBack(&node1)

	node2 := Node{2, nil, nil}
	list.PushBack(&node2)

	expected := 2
	if list.Size() != expected {
		t.Fatalf("bad push back: got %v, expected %d", list.Size(), expected)
	}

	list.Remove(&node2)
	expected = 1
	if list.Size() != expected {
		t.Fatalf("bad remove: got %v, expected %d", list.Size(), expected)
	}

	if list.Back() == nil {
		t.Fatalf("bad remove back: got %v", list.Back())
	}

	expected = 1
	if list.Back().Value != expected {
		t.Fatalf("bad remove back: got %v, expected %d", list.Back().Value, expected)
	}
}

func TestListRemoveMid(t *testing.T) {
	list := NewList()

	node1 := Node{1, nil, nil}
	list.PushBack(&node1)

	node2 := Node{2, nil, nil}
	list.PushBack(&node2)

	node3 := Node{3, nil, nil}
	list.PushBack(&node3)

	expected := 3
	if list.Size() != expected {
		t.Fatalf("bad push back: got %v, expected %d", list.Size(), expected)
	}

	list.Remove(&node2)
	expected = 2
	if list.Size() != expected {
		t.Fatalf("bad remove: got %v, expected %d", list.Size(), expected)
	}

	if list.Back() != &node3 {
		t.Fatalf("bad remove back: got %v", list.Back())
	}

	expected = 1
	if list.Node.Value != expected {
		t.Fatalf("bad remove front: got %v, expected %d", list.Back().Value, expected)
	}

	expected = 3
	if list.Back().Value != expected {
		t.Fatalf("bad remove back: got %v, expected %d", list.Back().Value, expected)
	}
}
