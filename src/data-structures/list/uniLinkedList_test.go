package uniLinkedList

import (
	"reflect"
	"testing"
)

func TestUniLinkedList(t *testing.T) {
	t.Run("it should Push", func(t *testing.T) {
		ull := CreateUniLinkedList[int]()
		const pushCount = 3
		expect := make([]int, 0, pushCount)
		for i := 0; i < pushCount; i++ {
			expect = append(expect, i)
			ull.Push(i)
		}

		got := ull.ToArray()

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect: %v, got: %v", expect, got)
		}
	})

	t.Run("it should Shift", func(t *testing.T) {
		ull := CreateUniLinkedList[int]()
		const shiftCount = 3
		expect := make([]int, 0, shiftCount)
		for i := 0; i < shiftCount; i++ {
			expect = append(expect, i)
			ull.Push(i)
		}
		got := make([]int, 0, shiftCount)

		for i := 0; i < shiftCount; i++ {
			value, err := ull.Shift()
			if err != nil {
				t.Errorf("shift error")
			}
			got = append(got, value)
		}

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect: %v, got: %v", expect, got)
		}
	})
}
