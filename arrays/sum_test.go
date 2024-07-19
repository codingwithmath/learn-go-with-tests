package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{4, 5})

	want := []int{3, 9}

	//reminder: Go do not let you equality operator with slices.
	if !reflect.DeepEqual(got, want) {
		//reflect is useful for seeing if ANY two variables are the same.
		//reflect is not type safe, so we can do silly comparissons like strings and slices
		t.Errorf("got %d want %d", got, want)
	}
}
