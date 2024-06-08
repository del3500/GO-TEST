package perimeter

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}
