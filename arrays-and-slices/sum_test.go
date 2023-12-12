package arrays_and_slices

import "testing"

func TestHello(t *testing.T) {

	t.Run("random size collection", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if want != got {
			t.Errorf("got %d, wanted %d given array %v", got, want, numbers)
		}
	})

}
