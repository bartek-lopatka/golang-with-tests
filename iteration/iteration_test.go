package iteration

import "testing"

func TestRepetition(t *testing.T) {
	t.Run("multiply character 5 times", func(t *testing.T) {
		char := Repeat("c")
		want := ("ccccc")

		if char != want {
			t.Errorf("got %q, want %q", char, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
