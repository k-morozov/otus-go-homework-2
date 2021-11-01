package homework_4

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type List struct {
	Node  *Node
	count int
}

func NewList() *List {
	return &List{}
}

func (l *List) Size() int {
	if nil == l {
		return 0
	}
	return l.count
}

func (l *List) Empty() bool {
	if nil == l {
		return true
	}
	return 0 == l.count
}

func (l *List) Front() *Node {
	if nil == l {
		return nil
	}
	return l.Node
}

func (l *List) Back() (lastNode *Node) {
	if nil == l {
		return
	}

	lastNode = l.Node

	if nil == lastNode {
		return
	}

	currentNode := l.Node

	for nil != currentNode {
		lastNode = currentNode
		currentNode = currentNode.Next
	}

	return
}

func (l *List) PushBack(value *Node) *Node {
	if nil == l {
		return nil
	}

	lastNode := l.Back()

	if nil != lastNode {
		lastNode.Next = value
		value.Prev = lastNode
	} else {
		l.Node = value
	}

	l.count++

	return value
}

func (l *List) Remove(node *Node) {
	if nil == l {
		return
	}

	prevNode := node.Prev
	nextNode := node.Next

	if nil != prevNode && nil != nextNode {
		prevNode.Next = nextNode
		nextNode.Prev = prevNode
	} else if nil != prevNode {
		prevNode.Next = nil
	} else if nil != nextNode {
		nextNode.Prev = nil
	} else {
		l.Node = nil
	}

	l.count--
}
