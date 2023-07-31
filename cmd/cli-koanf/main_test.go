package main

import (
	"os"
	"testing"
)

func TestNewKoanfFromEnvironment(t *testing.T) {
	want := "Jerry"
	os.Setenv("DEMO_PARENT1_CHILD1_NAME", want)

	k, err := NewKoanfFromEnvironment()
	if err != nil {
		t.Fatalf("failed to create new koanf value: %v", err)
	}

	t.Run("readKoanfString", func(t *testing.T) {
		got := k.String("parent1.child1.name")
		if got != want {
			t.Fatalf("want %v got %v", want, got)
		}
	})

	t.Run("unmarshal", func(t *testing.T) {
		var conf Conf
		if err := k.Unmarshal("", &conf); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		got := conf.Parent1.Child1.Name
		if got != want {
			t.Fatalf("want %v got %v", want, got)
		}
	})
}
