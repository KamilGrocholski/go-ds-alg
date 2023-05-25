package arrayMethods

import (
	"testing"
)

func TestArrayMethods(t *testing.T) {
	t.Run("it should Pop", func(t *testing.T) {
		array := []int{1}

		removed := Pop(&array)

		if removed != 1 {
			t.Errorf("pop return error")
		}

		if len(array) != 0 {
			t.Errorf("length error")
		}
	})

	t.Run("it should Shift", func(t *testing.T) {
		array := []int{1}

		removed := Shift(&array)

		if removed != 1 {
			t.Errorf("shift return error")
		}

		if len(array) != 0 {
			t.Errorf("length error")
		}
	})
}
