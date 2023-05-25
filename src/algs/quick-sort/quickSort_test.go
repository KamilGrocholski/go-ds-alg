package quickSort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	t.Run("it should sort", func(t *testing.T) {
		array := []int{6, 1, 5, 2, 0, 3, 4}
		expect := []int{0, 1, 2, 3, 4, 5, 6}

		QuickSort(&array, func(a, b int) bool {
			if a > b {
				return false
			}
			return true
		})

		if !reflect.DeepEqual(array, expect) {
			t.Errorf("expect: %v, got: %v", expect, array)
		}
	})
}
