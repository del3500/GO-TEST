package arrslice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	x := []int{5, 5, 5}
	y := []int{3, 3, 3}
	t.Run("collection of any size", func(t *testing.T) {

		got := Sum(x)
		want := 15

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("taking varying number of slices", func(t *testing.T) {
		got := SumAllOptimized(x, y)
		want := []int{15, 9}
		checkSum(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {

	t.Run("make the sum of all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{5, 2, 6})
		want := []int{5, 8}
		checkSum(t, got, want)
	})

	t.Run("safely sum empty slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 0}, []int{2, 1, 3, 10})
		want := []int{0, 14}
		checkSum(t, got, want)
	})
}

func checkSum(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkSum(b *testing.B) {
	x := []int{5, 5, 5}
	y := []int{3, 3, 3}
	for i := 0; i < b.N; i++ {
		SumAll(x, y)
	}
}

func BenchmarkSumAllBench1(b *testing.B) {
	x := []int{5, 5, 5}
	y := []int{3, 3, 3}
	for i := 0; i < b.N; i++ {
		SumAllBench1(x, y)
	}
}
func BenchmarkSumAllOptimized(b *testing.B) {
	x := []int{5, 5, 5}
	y := []int{3, 3, 3}
	for i := 0; i < b.N; i++ {
		SumAllOptimized(x, y)
	}
}
