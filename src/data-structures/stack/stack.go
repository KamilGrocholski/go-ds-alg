package stack

import "errors"

type UniNode[T any] struct {
	value T
	next  *UniNode[T]
}

type Stack[T any] struct {
	head *UniNode[T]
	tail *UniNode[T]
	size int
}

func CreateStack[T any]() *Stack[T] {
	return &Stack[T]{head: nil, tail: nil, size: 0}
}

func (stack *Stack[T]) Push(value T) {
	newNode := &UniNode[T]{value: value, next: nil}

	if stack.head == nil {
		stack.head = newNode
		stack.tail = newNode
	} else {
		newNode.next = stack.head
		stack.head = newNode
	}

	stack.size++
}

func (stack *Stack[T]) Pop() (T, error) {
	if stack.head == nil {
		var r T
		return r, errors.New("the stack is empty")
	}

	value := stack.head.value
	stack.head = stack.head.next

	return value, nil
}

func (stack *Stack[T]) Peek() (T, error) {
	if stack.head == nil {
		var r T
		return r, errors.New("the stack is empty")
	}

	return stack.head.value, nil
}

func (stack *Stack[T]) Clear() {
	stack.head = nil
	stack.tail = nil
	stack.size = 0
}
