package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	r := &Rectangle{
		Width:  5,
		Height: 5,
	}

	got := r.Perimeter()
	want := 20.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	r := &Rectangle{
		Width:  12.0,
		Height: 6.0,
	}

	got := r.Area()
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
