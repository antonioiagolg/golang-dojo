package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Width: 4, Height: 6}
	got := rectangle.Perimeter()
	want := 20.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}

}
func TestArea(t *testing.T) {

	areaTests := []struct {
        title   string
		shape   Shape
		hasArea float64
	}{
        {title: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
        {title: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
        {title: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.title, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v => got %g. want %g", tt, got, tt.hasArea)
			}
		})
	}
}
