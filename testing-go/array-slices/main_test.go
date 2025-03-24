package main

import (
	"slices"
	"testing"
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

	if !slices.Equal(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("Make sum of some slices", func(t *testing.T) {

		got := SumAllTails([]int{4, 5}, []int{0, 9})
		want := []int{5, 9}

		if !slices.Equal(got, want) {
			t.Errorf("Got %v, want %v", got, want)
		}
	})


	t.Run("Make sum of empty slices", func(t *testing.T) {

		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		if !slices.Equal(got, want) {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
}
