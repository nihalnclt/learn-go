package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %2.f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f want %2.f", got, want)
		}
	}

	// t.Run("rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{10.0, 10.0}
	// 	got := rectangle.Area()

	// 	checkArea(t, got, 100.0)
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{10.0}
	// 	got := circle.Area()

	// 	checkArea(t, got, 314.1592653589793)
	// })

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: &Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: &Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: &Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			checkArea(t, got, tt.hasArea)
		})
	}
}
