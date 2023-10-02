package lru

import "errors"

type Node struct {
	data string
	next *Node
	prev *Node
}

func NewNode(data string) *Node {
	return &Node{
		data: data,
		next: nil,
		prev: nil,
	}
}

type LinkedList struct {
	size int
	head *Node
	tail *Node
}

type ILinkedList interface {
	Get(index int) *Node
	InsertAtTail(data string) *Node
	DeleteFromHead() (*Node, error)
	MoveToTail(node *Node) error
}

func (ll *LinkedList) MoveToTail(node *Node) error {
	// The node itself is tail
	if node.next == nil {
		return nil
	}

	next := node.next // next is not nil
	prev := node.prev // in case of first node it can point to head

	next.prev = prev
	prev.next = next

	node.next = nil
	ll.tail.next = node
	ll.tail = node
	return nil
}

func (ll *LinkedList) Get(index int) *Node {

	ptr := ll.head.next
	for i := 0; i < index && ptr != nil; i++ {
		ptr = ptr.next
	}

	return ptr
}

func (ll *LinkedList) DeleteFromHead() (*Node, error) {
	node := ll.head.next
	// Linked list is empty
	if node == nil {
		return nil, errors.New("empty list")
	}

	ll.head.next = node.next
	if node.next != nil {
		node.next = ll.head
	}
	ll.size = ll.size - 1
	return node, nil
}

func (ll *LinkedList) InsertAtTail(key string) *Node {
	node := NewNode(key)
	ll.tail.next = node
	ll.tail = node
	ll.size = ll.size + 1
	return node
}

func NewLinkedList() *LinkedList {
	node := NewNode("")
	return &LinkedList{
		size: 0,
		head: node,
		tail: node,
	}
}
