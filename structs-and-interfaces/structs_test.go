package structs_and_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("counting perimeter of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("Got %.2f, want %.2f", got, want)
		}
	})

}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10.0}, 314.1592653589793},
		{Triangle{20.0, 10.0}, 100},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("Got %g, want %g", got, tt.want)
		}
	}
	//checkArea := func(t testing.TB, shape Shape, want float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if got != want {
	//		t.Errorf("Got %.2f, want %.2f", got, want)
	//	}
	//}
	//t.Run("counting area of a rectangle", func(t *testing.T) {
	//	rectangle := Rectangle{12.0, 6.0}
	//	checkArea(t, rectangle, 72)
	//})
	//
	//t.Run("counting area of a circle", func(t *testing.T) {
	//	circle := Circle{10.0}
	//	checkArea(t, circle, 314.1592653589793)
	//})

}
