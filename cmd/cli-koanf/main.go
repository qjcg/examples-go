package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

type Conf struct {
	Parent1 struct {
		Child1 struct {
			Name string
		}
	}
}

func NewKoanfFromEnvironment() (*koanf.Koanf, error) {
	k := koanf.New(".")

	if err := k.Load(env.Provider("DEMO_", ".", func(s string) string {
		envVar := strings.ToLower(strings.TrimPrefix(s, "DEMO_"))
		return strings.ReplaceAll(envVar, "_", ".")
	}), nil); err != nil {
		return k, err
	}

	return k, nil
}

func main() {
	k, err := NewKoanfFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	// This env var must be set for this to work.
	fmt.Printf("name from koanf value = %s\n", k.String("parent1.child1.name"))

	var conf Conf
	if err := k.Unmarshal("", &conf); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name from struct: %+v\n", conf)
}
