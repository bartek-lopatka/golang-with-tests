package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("random size collection", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if want != got {
			t.Errorf("got %d, wanted %d given array %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	t.Run("counting sum for variabe number of collections", func(t *testing.T) {
		got := SumAll([]int{2, 4}, []int{3, 6})
		want := []int{6, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})
}
func TestSumTails(t *testing.T) {
	t.Run("counting sum for tail sums of collections", func(t *testing.T) {
		got := SumTails([]int{2, 4}, []int{3, 6, 7})
		want := []int{4, 13}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
