package gofakeit

import (
	"fmt"
	"log"

	"github.com/brianvoe/gofakeit/v7"
)

type Mushroom string

func (m *Mushroom) Fake(f *gofakeit.Faker) (any, error) {
	return f.RandomString([]string{"cremini", "shiitake", "portobello"}), nil
}

// Custom function via implementation of the Fakeable interface.
func ExampleFakeable() {
	type Meal struct {
		Topping Mushroom
	}

	var m Meal
	err := gofakeit.Struct(&m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m.Topping)

	// Output: portobello
}
