package iteration

import (
	"testing"
    "fmt"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 6)
	expected := "aaaaaa"

	if got != expected {
		t.Errorf("Got %q, but want %q", got, expected)
	}
}

func ExampleRepeat() {
    got := Repeat("a", 5)
    fmt.Println(got)
    // Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
    for b.Loop() {
        Repeat("a", 1000)
    }
}
