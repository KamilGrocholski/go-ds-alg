package lru

import (
	"reflect"
	"testing"
)

func TestLru(t *testing.T) {
	t.Run("it should Update", func(t *testing.T) {
		const CAPACITY = 3
		lru := CreateLru[string, int](CAPACITY)
		users := map[string]int{"Kamil": 20, "Umibozu": 10}
		expect := make([]int, 0)

		for k, v := range users {
			expect = append(expect, v)
			lru.Update(k, v)
		}

		got := make([]int, 0)

		for k := range lru.lookup {
			recent, exists := lru.Get(k)

			if exists == false {
				t.Errorf("lookup error")
			}

			got = append(got, recent)
		}

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect: %v, got: %v", expect, got)
		}
	})
}
