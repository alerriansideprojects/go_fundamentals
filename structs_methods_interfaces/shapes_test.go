package shapes

import (
	"fmt"
	"testing"
)

func TestShapes(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()

		if got != want {
			t.Errorf("\nGot: %g\nWant: %g", got, want)
		}
	}

	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()

		got := shape.Perimeter()

		if got != want {
			t.Errorf("\nGot: %g\nWant: %g", got, want)
		}
	}

	t.Run("calculate perimeter of rectangles", func(t *testing.T) {
		rectangle := Rectangle{10, 10}
		checkPerimeter(t, rectangle, 40.0)
	})

	t.Run("calculate area of rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("area of circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}

func ExampleRectanglePerimeter() {
	rectangle := Rectangle{10.0, 10.0}
	perimeter := rectangle.Perimeter()

	fmt.Println(perimeter)
	//Output: 40
}

func ExampleRectangleArea() {
	rectangle := Rectangle{12.0, 6.0}
	area := rectangle.Area()

	fmt.Println(area)
	//Output: 72
}
