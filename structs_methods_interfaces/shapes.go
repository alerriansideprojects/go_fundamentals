package shapes

type Rectangle struct {
	Width float64
	Height float64
}

func rectanglePerimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func rectangleArea(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
