package queue

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("should Enqueue", func(t *testing.T) {
		queue := CreateQueue[int]()
		const enqueuesCount int = 3
		expect := make([]int, 0)

		for i := 0; i < enqueuesCount; i++ {
			expect = append(expect, i)
			queue.Enqueue((i))
			size := queue.Size()
			expectSize := i + 1
			if size != expectSize {
				t.Errorf("expect: size %v, got: size %v", expectSize, size)
			}
		}

		got := queue.ToArray()

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect: %v, got: %v", expect, got)
		}
	})

	t.Run("should Dequeue", func(t *testing.T) {
		queue := CreateQueue[int]()
		const enqueuesCount int = 3
		expect := make([]int, 0)

		for i := 0; i < enqueuesCount; i++ {
			expect = append(expect, i)
			queue.Enqueue(i)
		}

		got := make([]int, 0)

		for i := 0; i < enqueuesCount; i++ {
			value, err := queue.Dequeue()
			if err != nil {
				t.Errorf("an error from Dequeue")
			}
			size := queue.Size()
			expectSize := enqueuesCount - i - 1
			if size != expectSize {
				t.Errorf("expect: size %v, got: size %v", expectSize, size)
			}
			got = append(got, value)
		}

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect: %v, got: %v", expect, got)
		}
	})

	t.Run("should Clear", func(t *testing.T) {
		queue := CreateQueue[int]()

		queue.Enqueue(1)
		if queue.Size() != 1 {
			t.Errorf("enqueue size error")
		}
		queue.Clear()
		if queue.Size() != 0 {
			t.Errorf("clear size error")
		}
		if queue.head != nil {
			t.Errorf("head is not nil")
		}
		if queue.tail != nil {
			t.Errorf("tail is not nil")
		}
	})
}
