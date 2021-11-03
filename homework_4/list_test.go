package homework_4

import "testing"

func compare(lhs []int, rhs []int) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for i := 0; i < len(lhs); i++ {
		if lhs[i] != rhs[i] {
			return false
		}
	}

	return true
}

func TestNewListEmpty(t *testing.T) {
	list := NewList()

	const expected = 0

	if result := list.Size(); result != expected {
		t.Fatalf("bad size for new list: got %d, expected %d", result, expected)
	}

	if result := list.Empty(); result != true {
		t.Fatalf("bad empty for new list: got %v, expected %v", result, true)
	}

	if result := list.Front(); nil != result {
		t.Fatalf("bad front for new list: got %v, expected %v", result, nil)
	}

	if result := list.Back(); nil != result {
		t.Fatalf("bad back for new list: got %v, expected %v", result, nil)
	}
}

func TestListPushBackBasic(t *testing.T) {
	list := NewList()

	const expectedValue = 12
	newNode := list.PushBack(expectedValue)

	if newNode.Value != expectedValue {
		t.Fatalf("bad push back(value): got %v, expected %v", newNode.Value, expectedValue)
	}
	if newNode.Next != nil {
		t.Fatalf("bad push back(next): got %v, expected %v", newNode.Next, nil)
	}
	if newNode.Prev != nil {
		t.Fatalf("bad push back(prev): got %v, expected %v", newNode.Prev, nil)
	}

	const expectedSize = 1

	if expectedSize != list.Size() {
		t.Fatalf("bad push back: got %v, expected %v", list.Size(), expectedSize)
	}

	if result := list.Empty(); result != false {
		t.Fatalf("bad empty for new list: got %v, expected %v", result, false)
	}

	if result := list.Back(); nil == result {
		t.Fatalf("bad push back(back): got %v, expected %v", result, nil)
	}

	if result := list.Front(); nil == result {
		t.Fatalf("bad push back(front): got %v, expected %v", result, nil)
	}
}

func TestListPushFrontBasic(t *testing.T) {
	list := NewList()

	const expectedValue = 12
	newNode := list.PushFront(expectedValue)

	if newNode.Value != expectedValue {
		t.Fatalf("bad push front(value): got %v, expected %v", newNode.Value, expectedValue)
	}
	if newNode.Next != nil {
		t.Fatalf("bad push front(next): got %v, expected %v", newNode.Next, nil)
	}
	if newNode.Prev != nil {
		t.Fatalf("bad push front(prev): got %v, expected %v", newNode.Prev, nil)
	}

	const expectedSize = 1

	if expectedSize != list.Size() {
		t.Fatalf("bad push front: got %v, expected %v", list.Size(), expectedSize)
	}

	if result := list.Empty(); result != false {
		t.Fatalf("bad empty for new list: got %v, expected %v", result, false)
	}

	if result := list.Back(); nil == result {
		t.Fatalf("bad push front(back): got %v, expected %v", result, nil)
	}

	if result := list.Front(); nil == result {
		t.Fatalf("bad push front(front): got %v, expected %v", result, nil)
	}
}

func TestListSomePushBack(t *testing.T) {
	list := NewList()

	const expectedValueFront = 1
	const expectedValueMid = 2
	const expectedValueBack = 3

	list.PushBack(expectedValueFront)
	list.PushBack(expectedValueMid)
	list.PushBack(expectedValueBack)

	const expectedSize = 3

	if expectedSize != list.Size() {
		t.Fatalf("bad some push back(size): got %v, expected %v", list.Size(), expectedSize)
	}

	if result := list.Front(); result == nil {
		t.Fatalf("bad some push back(front): got %v", result)
	}

	if result := list.Front(); expectedValueFront != result.Value {
		t.Fatalf("bad some push back(front value): got %v, expected %v", result, expectedValueFront)
	}

	if result := list.Front().Prev; nil != result {
		t.Fatalf("bad some push back(front prev): got %v, expected %v", result, nil)
	}

	if result := list.Front().Next; nil == result {
		t.Fatalf("bad some push back(front next): got %v, expected %v", result, nil)
	}

	if result := list.Back(); expectedValueBack != result.Value {
		t.Fatalf("bad some push back(back value): got %v, expected %v", result, expectedValueBack)
	}

	if result := list.Back().Next; nil != result {
		t.Fatalf("bad some push back(back next): got %v, expected %v", result, nil)
	}

	if result := list.Back().Prev; nil == result {
		t.Fatalf("bad some push back(back prev): got %v, expected %v", result, expectedValueBack)
	}

	if list.Front().Next != list.Back().Prev {
		t.Fatalf("bad some push back(mid): got %v, expected %v", list.Front().Next, list.Back().Prev)
	}

	if result := list.Back().Prev; expectedValueMid != result.Value {
		t.Fatalf("bad some push back(back prev value): got %v, expected %v", result, expectedValueMid)
	}

	if result := list.Front().Next; expectedValueMid != result.Value {
		t.Fatalf("bad some push back(front next value): got %v, expected %v", result, expectedValueMid)
	}

	expected := []int{expectedValueFront, expectedValueMid, expectedValueBack}
	results := make([]int, 0, list.Size())
	for node := list.Front(); node != nil; node = node.Next {
		results = append(results, node.Value.(int))
	}

	if !compare(expected, results) {
		t.Fatalf("bad some push back: got %v, expected %v", results, expected)
	}
}

func TestListSomePushFront(t *testing.T) {
	list := NewList()

	const expectedValueFront = 1
	const expectedValueMid = 2
	const expectedValueBack = 3

	list.PushFront(expectedValueBack)
	list.PushFront(expectedValueMid)
	list.PushFront(expectedValueFront)

	const expectedSize = 3

	if expectedSize != list.Size() {
		t.Fatalf("bad some push front(size): got %v, expected %v", list.Size(), expectedSize)
	}

	if result := list.Front(); result == nil {
		t.Fatalf("bad some push front(front): got %v", result)
	}

	if result := list.Front(); expectedValueFront != result.Value {
		t.Fatalf("bad some push front(front value): got %v, expected %v", result, expectedValueFront)
	}

	if result := list.Front().Prev; nil != result {
		t.Fatalf("bad some push front(front prev): got %v, expected %v", result, nil)
	}

	if result := list.Front().Next; nil == result {
		t.Fatalf("bad some push front(front next): got %v, expected %v", result, nil)
	}

	if result := list.Back(); expectedValueBack != result.Value {
		t.Fatalf("bad some push front(back value): got %v, expected %v", result, expectedValueBack)
	}

	if result := list.Back().Next; nil != result {
		t.Fatalf("bad some push front(back next): got %v, expected %v", result, nil)
	}

	if result := list.Back().Prev; nil == result {
		t.Fatalf("bad some push front(back prev): got %v, expected %v", result, expectedValueBack)
	}

	if list.Front().Next != list.Back().Prev {
		t.Fatalf("bad some push front(mid): got %v, expected %v", list.Front().Next, list.Back().Prev)
	}

	if result := list.Back().Prev; expectedValueMid != result.Value {
		t.Fatalf("bad some push front(back prev value): got %v, expected %v", result, expectedValueMid)
	}

	if result := list.Front().Next; expectedValueMid != result.Value {
		t.Fatalf("bad some push front(front next value): got %v, expected %v", result, expectedValueMid)
	}

	expected := []int{expectedValueFront, expectedValueMid, expectedValueBack}
	results := make([]int, 0, list.Size())
	for node := list.Front(); node != nil; node = node.Next {
		results = append(results, node.Value.(int))
	}

	if !compare(expected, results) {
		t.Fatalf("bad some push front: got %v, expected %v", results, expected)
	}
}

func TestListRemoveHead(t *testing.T) {
	list := NewList()

	node := list.PushBack(1)

	expected := 1
	if list.Size() != expected {
		t.Fatalf("bad push back: got %v, expected %d", list.Size(), expected)
	}

	list.Remove(node)
	expected = 0
	if list.Size() != expected {
		t.Fatalf("bad remove: got %v, expected %d", list.Size(), expected)
	}
}

func TestListRemoveTail(t *testing.T) {
	list := NewList()

	list.PushBack(1)
	node2 := list.PushBack(2)

	expected := 2
	if list.Size() != expected {
		t.Fatalf("bad push back: got %v, expected %d", list.Size(), expected)
	}

	list.Remove(node2)
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

	list.PushBack(1)
	node2 := list.PushBack(2)
	node3 := list.PushBack(3)

	expected := 3
	if list.Size() != expected {
		t.Fatalf("bad push back: got %v, expected %d", list.Size(), expected)
	}

	list.Remove(node2)
	expected = 2
	if list.Size() != expected {
		t.Fatalf("bad remove: got %v, expected %d", list.Size(), expected)
	}

	if list.Back() != node3 {
		t.Fatalf("bad remove back: got %v", list.Back())
	}

	expected = 1
	if list.Front().Value != expected {
		t.Fatalf("bad remove front: got %v, expected %d", list.Back().Value, expected)
	}

	expected = 3
	if list.Back().Value != expected {
		t.Fatalf("bad remove back: got %v, expected %d", list.Back().Value, expected)
	}
}

func TestList_MoveToFrontSize1(t *testing.T) {
	list := NewList()

	const expectedFirst = 100

	moveNode := list.PushBack(expectedFirst)

	list.MoveToFront(moveNode)

	if result := list.Front(); moveNode != result {
		t.Fatalf("bad some move to front: got %v, expected %v", result, moveNode)
	}

	if result := list.Front().Value; expectedFirst != result {
		t.Fatalf("bad some move to front: got %v, expected %v", result, expectedFirst)
	}

	if result := list.Back().Value; expectedFirst != result {
		t.Fatalf("bad some move to front: got %v, expected %v", result, expectedFirst)
	}

	if result := list.Front().Prev; nil != result {
		t.Fatalf("bad some move to front: got %v, expected %v", result, nil)
	}
}

func TestList_MoveToFrontSize2Begin(t *testing.T) {
	list := NewList()

	const expectedFirst = 100
	const expectedSecond = 200

	moveNode := list.PushBack(expectedFirst)
	list.PushBack(expectedSecond)

	list.MoveToFront(moveNode)

	if result := list.Front(); moveNode != result {
		t.Fatalf("bad some move to front: got %v, expected %v", result, moveNode)
	}

	if result := list.Front().Value; expectedFirst != result {
		t.Fatalf("bad some move to front: got %v, expected %v", result, expectedFirst)
	}

	if result := list.Back().Value; expectedSecond != result {
		t.Fatalf("bad some move to front: got %v, expected %v", result, expectedSecond)
	}

	if result := list.Back().Next; nil != result {
		t.Fatalf("bad some move to front: got %v, expected %v", result, nil)
	}
}

func TestList_MoveToFrontSize2End(t *testing.T) {
	list := NewList()

	const expectedFirst = 100
	const expectedSecond = 200

	list.PushBack(expectedFirst)
	moveNode := list.PushBack(expectedSecond)

	list.MoveToFront(moveNode)

	if result := list.Front(); moveNode != result {
		t.Fatalf("bad some move to front: got %v, expected %v", result, moveNode)
	}

	if result := list.Front().Value; expectedSecond != result {
		t.Fatalf("bad some move to front: got %v, expected %v", result, expectedSecond)
	}

	if result := list.Back().Value; expectedFirst != result {
		t.Fatalf("bad some move to front: got %v, expected %v", result, expectedFirst)
	}

	if result := list.Front().Prev; nil != result {
		t.Fatalf("bad some move to front: got %v, expected %v", result, nil)
	}

	if result := list.Back().Next; nil != result {
		t.Fatalf("bad some move to front: got %v, expected %v", result, nil)
	}
}

func TestList_MoveToFrontSize3Mid(t *testing.T) {
	list := NewList()

	const expectedFirst = 100
	const expectedSecond = 200
	const expectedThird = 300

	list.PushBack(expectedFirst)
	moveNode := list.PushBack(expectedSecond)
	list.PushBack(expectedThird)
	list.MoveToFront(moveNode)

	if result := list.Front(); moveNode != result {
		t.Fatalf("bad some move to front: got %v, expected %v", result, moveNode)
	}

	if result := list.Front().Value; expectedSecond != result {
		t.Fatalf("bad some move to front: got %v, expected %v", result, expectedSecond)
	}

	if result := list.Back().Value; expectedThird != result {
		t.Fatalf("bad some move to front: got %v, expected %v", result, expectedThird)
	}

	if result := list.Front().Value; expectedSecond != result {
		t.Fatalf("bad some move to front: got %v, expected %v", result, expectedSecond)
	}

	if result := list.Front().Next.Value; expectedFirst != result {
		t.Fatalf("bad some move to front: got %v, expected %v", result, expectedFirst)
	}

	if result := list.Back().Prev.Value; expectedFirst != result {
		t.Fatalf("bad some move to front: got %v, expected %v", result, expectedFirst)
	}

	if result := list.Front().Prev; nil != result {
		t.Fatalf("bad some move to front: got %v, expected %v", result, nil)
	}
}
