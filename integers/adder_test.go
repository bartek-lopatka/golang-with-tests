package integers

import "testing"

func testAdder(t *testing.T) {
	t.Run("adding two integers", func(t *testing.T) {
		sum := Add(2, 4)
		want := 6

		if sum != want {
			t.Errorf("sum %d, wanted %d", sum, want)
		}
	})
}
