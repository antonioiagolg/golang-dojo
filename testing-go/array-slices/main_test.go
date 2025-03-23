package main

import (
    "testing"
    "reflect"
)

func TestSum(t *testing.T) {

	t.Run("Collection of any size", func(t *testing.T) {

		nums := []int{1, 2, 3}
		got := Sum(nums)
		want := 6

		if got != want {
			t.Errorf("Got %d, want %d", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %d, want %d", got, want)
	}
}
