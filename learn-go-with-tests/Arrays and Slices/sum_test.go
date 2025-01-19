package main

import "testing"

func TestSum(t *testing.T) {
	// t.Run("collection of 5 numbers", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3, 4, 5}
	//
	// 	got := Sum(numbers)
	// 	want := 15
	//
	// 	if got != want {
	// 		t.Errorf("got %d want %d; numbers: %v", got, want, numbers)
	// 	}
	// })

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{10, 20, 40}

		got := Sum(numbers)
		want := 70

		if got != want {
			t.Errorf("got %d want %d; numbers: %v", got, want, numbers)
		}
	})
}
