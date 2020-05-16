package shapes

import (
	"fmt"
	"testing"
)

func TestShapes(t *testing.T) {
	assertCorrectCalculation := func(t *testing.T, got, want float64) {
		t.Helper()

		if got != want {
			t.Errorf("\nGot: %.2f\nWant: %.2f", got, want)
		}
	}

	t.Run("calculate perimeter of rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectanglePerimeter(rectangle)
		want := 40.0

		assertCorrectCalculation(t, got, want)
	})

	t.Run("calculate area of rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := rectangleArea(rectangle)
		want := 72.0

		assertCorrectCalculation(t, got, want)
	})
}

func ExampleRectanglePerimeter() {
	rectangle := Rectangle{10.0, 10.0}
	perimeter := rectanglePerimeter(rectangle)

	fmt.Println(perimeter)
	//Output: 40
}

func ExampleRectangleArea() {
	rectangle := Rectangle{12.0, 6.0}
	area := rectangleArea(rectangle)

	fmt.Println(area)
	//Output: 72
}

