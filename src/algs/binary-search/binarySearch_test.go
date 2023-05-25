package binarySearch

import "testing"

func TestBinarySearch(t *testing.T) {
	t.Run("it should find", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5}
		expect := 1

		got := BinarySearch(&array, 2, func(a, b int) int {
			if a < b {
				return -1
			}
			if a > b {
				return 1
			}
			return 0
		})

		if got != expect {
			t.Errorf("expect: %v, got: %v", expect, got)
		}
	})
}
