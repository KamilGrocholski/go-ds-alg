package stack

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("should Push", func(t *testing.T) {
		stack := CreateStack[int]()
		const pushCount = 3
		expect := make([]int, 0, pushCount)

		for i := 0; i < pushCount; i++ {
			expect = append(expect, i)
			stack.Push(i)
		}

		got := make([]int, 0, pushCount)

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect: %v, got: %v", expect, got)
		}
	})

	t.Run("should Pop", func(t *testing.T) {
		stack := CreateStack[int]()
		const popCount = 3
		expect := make([]int, 0, popCount)

		for i := 0; i < popCount; i++ {
			expect = append(expect, i)
			stack.Push(i)
		}

		got := make([]int, 0, popCount)

		for i := 0; i < popCount; i++ {
			value, err := stack.Pop()
			if err != nil {
				t.Errorf("pop error")
			}
			got = append(got, value)
		}

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect: %v, got: %v", expect, got)
		}
	})
}
