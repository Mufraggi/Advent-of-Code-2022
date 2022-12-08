package main

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	len  int
}

func (l *LinkedList[T]) PushBack(val T) {
	n := Node[T]{}
	n.value = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

func (l *LinkedList[T]) UpdateAt(pos int, val T) *Node[T] {
	ptr := l.head
	if pos < 0 {
		return ptr
	}
	if pos > l.len-1 {
		return nil
	}
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	ptr.value = val
	return ptr
}

func (l LinkedList[T]) printList() {
	if l.len == 0 {
		fmt.Println("no node")
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println(ptr.value)
		ptr = ptr.next
	}

}
