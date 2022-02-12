package tested

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Shape interface {
	Area() float64
	Perimeter() float64
}
