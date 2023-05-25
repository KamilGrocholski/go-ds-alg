package minHeap

type IsSmallerFn[T any] func(a, b T) bool

type MinHeap[T any] struct {
	storage   []T
	IsSmaller IsSmallerFn[T]
}

func Swap[T any](array []T, a int, b int) {
	temp := array[a]
	array[a] = array[b]
	array[b] = temp
}

func CreateMinHeap[T any](IsSmaller IsSmallerFn[T]) *MinHeap[T] {
	return &MinHeap[T]{storage: make([]T, 0), IsSmaller: IsSmaller}
}

func (heap *MinHeap[T]) Push(value T) {
	heap.storage = append(heap.storage, value)
}

func (heap *MinHeap[T]) Pop() T {
	removed := heap.storage[len(heap.storage)-1]

	if len(heap.storage) == 0 {
		var nothing T
		return nothing
	}

	(heap.storage) = (heap.storage)[:len(heap.storage)-1]

	return removed
}

func HeapifyUp[T any](heap *MinHeap[T], index int) {
	if index == 0 {
		return
	}

	parentIndex := GetParentIndex(index)
	parentValue := heap.storage[parentIndex]
	childValue := heap.storage[index]
	isChildSmaller := heap.IsSmaller(parentValue, childValue)

	if isChildSmaller == false {
		Swap(heap.storage, index, parentIndex)
		HeapifyUp(heap, parentIndex)
	}
}

func HeapifyDown[T any](heap *MinHeap[T], index int) {
	leftChildIndex := GetLeftChildIndex(index)
	rightChildIndex := GetRightChildIndex(index)

	storageSize := len(heap.storage)

	if index >= storageSize || leftChildIndex >= storageSize {
		return
	}

	leftChildValue := heap.storage[leftChildIndex]
	rightChildValue := heap.storage[rightChildIndex]
	currentValue := heap.storage[index]

	isLeftChildSmaller := heap.IsSmaller(leftChildValue, rightChildValue)

	if isLeftChildSmaller == false && heap.IsSmaller(currentValue, rightChildValue) == false {
		Swap(heap.storage, index, rightChildIndex)
		HeapifyDown(heap, rightChildIndex)
	} else if isLeftChildSmaller == true && heap.IsSmaller(currentValue, leftChildValue) == true {
		Swap(heap.storage, index, leftChildIndex)
		HeapifyDown(heap, leftChildIndex)
	}
}

func GetParentIndex(index int) int {
	return (index - 1) / 2
}

func GetLeftChildIndex(index int) int {
	return (index * 2) + 1
}

func GetRightChildIndex(index int) int {
	return (index * 2) + 2
}
