package iteration

import (
	"strings"
	"testing"
)

func TestIteration(t *testing.T) {

	t.Run("repeat a letter", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := strings.Repeat("a", 5)

		assertRepeated(t, repeated, expected)
	})

	t.Run("multiple letters", func(t *testing.T) {
		repeated := Repeat("abc", 3)
		expected := strings.Repeat("abc", 3)

		assertRepeated(t, repeated, expected)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func assertRepeated(t testing.TB, r, e string) {
	t.Helper()

	if r != e {
		t.Errorf("repeated %s expected %s", r, e)
	}
}
