package binaryTree

import (
	"reflect"
	"testing"
)

func Compare(a int, b int) Comparison {
	if a > b {
		return BIGGER
	}
	if a < b {
		return SMALLER
	}
	return EQUAL
}

func TestBinaryTre(t *testing.T) {
	t.Run("it should Push", func(t *testing.T) {
		tree := CreateBinaryTree[int](Compare)
		const pushCount = 3
		expect := make([]int, 0, pushCount)

		for i := 0; i < pushCount; i++ {
			expect = append(expect, i)
			tree.Push(i)
		}

		got := tree.ToArray()

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect: %v, got: %v", expect, got)
		}
	})
}
