package structs_and_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("counting perimeter of rectangle", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0

		if got != want {
			t.Errorf("Got %.2f, want %.2f", got, want)
		}
	})

}
