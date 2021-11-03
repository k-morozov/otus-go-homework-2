package homework_4

type ForwardList interface {
	Size() int
	Empty() bool
	Front() *Node
	Back() (lastNode *Node)
	PushFront(value interface{}) *Node
	PushBack(value interface{}) *Node
	Remove(node *Node)
}

type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

type list struct {
	Node  *Node
	count int
}

func NewList() ForwardList {
	return &list{}
}

func (l *list) Size() int {
	if nil == l {
		return 0
	}
	return l.count
}

func (l *list) Empty() bool {
	if nil == l {
		return true
	}
	return 0 == l.count
}

func (l *list) Front() *Node {
	if nil == l {
		return nil
	}
	return l.Node
}

func (l *list) Back() (lastNode *Node) {
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

func (l *list) PushFront(value interface{}) *Node {
	if nil == l {
		return nil
	}

	temp := &Node{Value: value}
	if nil != l.Node {
		l.Node.Prev = temp
	}
	temp.Next = l.Node
	l.Node = temp

	l.count++
	return l.Node
}

func (l *list) PushBack(value interface{}) *Node {
	if nil == l {
		return nil
	}

	lastNode := l.Back()

	if nil != lastNode {
		temp := &Node{Value: value}
		lastNode.Next = temp
		temp.Prev = lastNode
	} else {
		l.Node = &Node{Value: value}
	}

	l.count++

	return l.Back()
}

func (l *list) Remove(node *Node) {
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
