package greet_test

import (
	"fmt"
	"testing"

	"github.com/qjcg/go-examples/internal/greet"
)

func TestGreet(t *testing.T) {
	testCases := []struct {
		description string
		name        string
		want        string
	}{
		{"empty", "", "Hello!"},
		{"basic1", "Jerry", "Hello, Jerry!"},
		{"basic2", "Newman", "Hello, Newman!"},
		{"basic3", "Elaine", "Hello, Elaine!"},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := greet.Greet(tc.name)
			if got != tc.want {
				t.Fatalf("Want: %s, Got: %s\n", tc.want, got)
			}

		})

	}
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		greet.Greet("Jerry")
	}

}

func ExampleGreet() {
	greetEmpty := greet.Greet("")
	fmt.Println(greetEmpty)

	greetLeo := greet.Greet("Leo")
	fmt.Println(greetLeo)

	// Output:
	// Hello!
	// Hello, Leo!
}
