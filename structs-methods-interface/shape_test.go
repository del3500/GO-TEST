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

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		r := Rectangle{
			Width:  12.0,
			Height: 6.0,
		}
		checkArea(t, r, 72.0)
	})
	t.Run("circles", func(t *testing.T) {
		c := Circle{10}
		checkArea(t, c, 314.1592653589793)
	})
}
