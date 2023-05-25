package minHeap

import (
	"reflect"
	"testing"
)

func IsSmaller(a, b int) bool {
	if a < b {
		return true
	}

	return false
}

func TestMinHeap(t *testing.T) {
	t.Run("it should Push", func(t *testing.T) {
		heap := CreateMinHeap[int](IsSmaller)
		const pushCount = 3
		expect := make([]int, 0)

		for i := 0; i < pushCount; i++ {
			heap.Push(i)
			expect = append(expect, i)
		}

		got := heap.storage

		if heap.storage[0] > heap.storage[1] || heap.storage[0] > heap.storage[2] {
			t.Errorf("the heap is not valid")
		}

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect: %v, got: %v", expect, got)
		}
	})

	t.Run("it should Pop", func(t *testing.T) {
		heap := CreateMinHeap[int](IsSmaller)
		expect := []int{1, 2, 3}

		for _, value := range expect {
			heap.Push(value)
		}

		got := make([]int, len(expect))

		for index := range expect {
			removed := heap.Pop()
			got[2-index] = removed
		}

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect: %v, got: %v", expect, got)
		}
	})
}
