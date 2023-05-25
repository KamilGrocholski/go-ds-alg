package queue

import "errors"

type UniNode[T any] struct {
	value T
	next  *UniNode[T]
}

type Queue[T any] struct {
	head *UniNode[T]
	tail *UniNode[T]
	size int
}

func CreateQueue[T any]() *Queue[T] {
	return &Queue[T]{head: nil, tail: nil, size: 0}
}

func (queue *Queue[T]) Enqueue(value T) {
	newNode := &UniNode[T]{value: value, next: nil}

	if queue.tail == nil {
		queue.tail = newNode
		queue.head = newNode
	} else {
		queue.tail.next = newNode
		queue.tail = newNode
	}

	queue.size++
}

func (queue *Queue[T]) Dequeue() (T, error) {
	if queue.head == nil {
		var r T
		return r, errors.New("the queue is empty")
	}

	value := queue.head.value
	queue.head = queue.head.next

	queue.size--

	return value, nil
}

func (queue *Queue[T]) Size() int {
	return queue.size
}

func (queue *Queue[T]) ToArray() []T {
	array := make([]T, 0, int(queue.Size()))

	node := queue.head

	for node != nil {
		array = append(array, node.value)
		node = node.next
	}

	return array
}

func (queue *Queue[T]) IsEmpty() bool {
	return queue.size == 0
}

func (queue *Queue[T]) Clear() {
	queue.head = nil
	queue.tail = nil
	queue.size = 0
}
