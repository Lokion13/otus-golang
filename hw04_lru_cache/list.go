package hw04lrucache

import (
	"fmt"
	"log"
	"strconv"
)

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Key   Key
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

func NewListItem(i interface{}) *ListItem {
	var res ListItem
	switch v := i.(type) {
	case int:
		strKey := Key(strconv.Itoa(i.(int)))
		res = ListItem{strKey, v, nil, nil}
	case string:
		res = ListItem{Key(v), v, nil, nil}
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
	return &res
}

type ListD struct {
	length int
	head   *ListItem
	tail   *ListItem
}

func NewList() *ListD {
	return &ListD{}
}

func (l *ListD) Len() int {
	return l.length
}

func (l *ListD) Front() *ListItem {
	return l.head
}

func (l *ListD) Back() *ListItem {
	return l.tail
}

func (l *ListD) PushFront(v interface{}) *ListItem {
	node := NewListItem(v)
	if l.length <= 0 {
		l.head = node
		l.tail = node
		l.length = 1
		return node
	}
	if l.head == nil {
		log.Fatal("head is nil -- ", l)
	}
	l.head.Prev = node
	node.Next = l.head
	node.Prev = nil
	l.head = node
	l.length++
	return node
}

func (l *ListD) PushBack(v interface{}) *ListItem {
	node := NewListItem(v)
	if l.length <= 0 {
		l.head = node
		l.tail = node
		l.length = 1
		return node
	}
	l.tail.Next = node
	node.Prev = l.tail
	l.tail = node
	l.length++
	return node
}

func (l *ListD) Remove(i *ListItem) {
	if i == nil || l.length == 0 {
		log.Fatal("cannot remove - node == nil")
	}

	if i.Prev == nil {
		l.head = i.Next
		i.Next.Prev = nil
	}
	if i.Next == nil {
		l.tail = i.Prev
		i.Prev.Next = nil
	}
	if i.Prev != nil && i.Next != nil {
		i.Next.Prev = i.Prev
		i.Prev.Next = i.Next
	}
	l.length--
}

func (l *ListD) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}
	prev := i.Prev
	next := i.Next
	if i.Next == nil {
		prev.Next = nil
		l.tail = prev
	} else {
		prev.Next = next
		next.Prev = prev
	}
	l.head.Prev = i
	i.Next = l.head
	i.Prev = nil
	l.head = i
}
