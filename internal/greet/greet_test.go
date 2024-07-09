package greet_test

import (
	"fmt"
	"testing"

	"github.com/qjcg/examples-go/internal/greet"
)

func TestGreet(t *testing.T) {
	f := func(name, want string) {
		t.Helper()

		got := greet.Greet(name)
		if got != want {
			t.Fatalf("want: %q, got %q", want, got)
		}
	}

	// Empty
	f("", "Hello!")

	// Names
	f("Newman", "Hello, Newman!")
	f("Elaine", "Hello, Elaine!")
	f("Kramer", "Hello, Kramer!")
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
