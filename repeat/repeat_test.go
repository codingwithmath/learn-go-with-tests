package repeat

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	result := Repeat("a", 6)
	expected := "aaaaaa"

	if result != expected {
		t.Errorf("got %q want %q", result, expected)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 4)
	fmt.Println(repeated)

	//Output: aaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}
