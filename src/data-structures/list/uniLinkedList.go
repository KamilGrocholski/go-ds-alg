package uniLinkedList

import (
	"errors"
)

type UniNode[T any] struct {
	value T
	next  *UniNode[T]
}

type UniLinkedList[T any] struct {
	head *UniNode[T]
	tail *UniNode[T]
	size int
}

func CreateUniLinkedList[T any]() *UniLinkedList[T] {
	return &UniLinkedList[T]{head: nil, tail: nil, size: 0}
}

func (list *UniLinkedList[T]) Push(value T) {
	newNode := &UniNode[T]{value: value, next: nil}

	list.size++

	if list.tail == nil {
		list.tail = newNode
		list.head = newNode
	} else {
		list.tail.next = newNode
		list.tail = newNode
	}
}

func (list *UniLinkedList[T]) Unshift(value T) {
	newNode := &UniNode[T]{value: value, next: nil}

	list.size++

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head = newNode
	}
}

func (list *UniLinkedList[T]) Shift() (T, error) {
	if list.head == nil {
		var r T
		return r, errors.New("the list is empty")
	}

	value := list.head.value

	newHead := list.head.next
	list.head = nil
	list.head = newHead
	list.size--

	return value, nil
}

func (list *UniLinkedList[T]) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

func (list *UniLinkedList[T]) ForEach(callback func(value T)) {
	node := list.head

	for node != nil {
		callback(node.value)
		node = node.next
	}
}

func (list *UniLinkedList[T]) ToArray() []T {
	array := make([]T, 0, int(list.size))

	node := list.head

	for node != nil {
		array = append(array, node.value)
		node = node.next
	}

	return array
}

func (list *UniLinkedList[T]) IsEmpty() bool {
	return list.size == 0
}

func (list *UniLinkedList[T]) Size() int {
	return list.size
}

type UniLinkedListIterator[T any] struct {
	nextNode *UniNode[T]
}

func (iterator *UniLinkedListIterator[T]) GetElement() *UniNode[T] {
	return iterator.nextNode
}

func (iterator *UniLinkedListIterator[T]) Next() *UniNode[T] {
	if iterator == nil {
		return nil
	}

	iterator.nextNode = iterator.nextNode.next
	return iterator.nextNode
}

func (list *UniLinkedList[T]) GetIterator() UniLinkedListIterator[T] {
	return UniLinkedListIterator[T]{nextNode: list.head}
}
