package hw04lrucache

import "fmt"

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	Print()
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len  int
	head *ListItem
	tail *ListItem
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem {
	return l.head
}

func (l list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	node := &ListItem{v, l.head, nil}
	if l.head == nil {
		l.tail = node
	} else {
		l.head.Prev = node
	}
	l.len++
	l.head = node
	return l.head
}

func (l *list) PushBack(v interface{}) *ListItem {
	node := &ListItem{v, nil, l.tail}
	if l.tail == nil {
		l.head = node
	} else {
		l.tail.Next = node
	}
	l.len++
	l.tail = node
	return l.tail
}

func (l *list) Remove(i *ListItem) {
	if i.Prev == nil {
		l.head = i.Next
	} else {
		i.Prev.Next = i.Next
	}
	if i.Next == nil {
		l.tail = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.PushFront(i.Value)
	l.Remove(i)
}

func (l list) Print() {
	node := l.head
	fmt.Println("Len: ", l.len)
	fmt.Print("List: nil")
	for node != nil {
		fmt.Print(" <-> ", node.Value)
		node = node.Next
	}
	fmt.Println(" <-> nil")
}

func NewList() List {
	return new(list)
}
