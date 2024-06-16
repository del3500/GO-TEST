package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {

	buffer := bytes.Buffer{}
	Greet(&buffer, "Del")

	got := buffer.String()
	want := "Hello, Del"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
