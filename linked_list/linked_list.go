package main

type Node struct {
	data interface{}
	next *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) InsertTail(data interface{}) {
	node := &Node{
		data: data,
		next: nil,
	}
	if l.head == nil {
		l.head = node
		return
	}

	temp := l.head

	for temp.next != nil {
		temp = temp.next
	}

	temp.next = node
}

func (l *LinkedList) InsertHead(data interface{}) {
	node := &Node{
		data: data,
		next: nil,
	}
	if l.head == nil {
		l.head = node
		return
	}

	node.next = l.head

	l.head = node
}
