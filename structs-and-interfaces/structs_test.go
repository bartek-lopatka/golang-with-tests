package structs_and_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("counting perimeter of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("Got %.2f, desiredArea %.2f", got, want)
		}
	})

}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name        string
		shape       Shape
		desiredArea float64
	}{
		{"Rectangle", Rectangle{12.0, 6.0}, 72.0},
		{"Circle", Circle{10.0}, 314.1592653589793},
		{"Triangle", Triangle{20.0, 10.0}, 100},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.desiredArea {
				t.Errorf("Shape %#v got %g, desiredArea is %g", tt.shape, got, tt.desiredArea)
			}
		})
	}
	//checkArea := func(t testing.TB, shape Shape, desiredArea float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if got != desiredArea {
	//		t.Errorf("Got %.2f, desiredArea %.2f", got, desiredArea)
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
