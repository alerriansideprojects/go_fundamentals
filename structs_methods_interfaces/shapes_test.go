package shapes

import (
	"testing"
)

func TestArea(t *testing.T) {
	areaTests := []struct {
		name string
		shape Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v\nGot: %g\nWant: %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name string
		shape Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{10, 10}, hasPerimeter: 40.0},
		{name: "Circle", shape: Circle{20}, hasPerimeter: 125.66370614359172},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.hasPerimeter {
				t.Errorf("%#v\nGot: %g\nWant: %g", tt.shape, got, tt.hasPerimeter)
			}
		})
	}
}
