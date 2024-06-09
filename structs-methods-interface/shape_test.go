package perimeter

import "testing"

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t testing.TB, r Rectangle, want float64) {
		t.Helper()
		got := r.Perimeter()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangle perimeter", func(t *testing.T) {
		r := Rectangle{
			Width:  5,
			Height: 5,
		}
		checkPerimeter(t, r, 20.0)
	})
}

func TestArea(t *testing.T) {

	areaTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}

}
