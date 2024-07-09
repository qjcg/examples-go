// Package greet provides greeting-related functions.
package greet

import "fmt"

// Greet returns a friendly greeting.
func Greet(name string) string {
	if name == "" {
		return "Hello!"
	}

	return fmt.Sprintf("Hello, %s!", name)
}
