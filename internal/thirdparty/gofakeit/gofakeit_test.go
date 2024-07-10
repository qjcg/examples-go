package gofakeit

import (
	"fmt"
	"log"

	"github.com/brianvoe/gofakeit/v7"
)

func init() {
	// Using a constant seed makes the output of Examples deterministic.
	err := gofakeit.Seed(123)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleStruct() {
	type User struct {
		Id    int    `fake:"{number:1,100}"`
		Name  string `fake:"{firstname}"`
		Email string `fake:"{email}"`
	}

	var u User
	err := gofakeit.Struct(&u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u.Id)
	fmt.Println(u.Name)
	fmt.Println(u.Email)

	// Output:
	// 86
	// Harold
	// francesvonrueden@heathcote.net
}

func ExampleFunc() string {
	gofakeit.AddFuncLookup("friendname", gofakeit.Info{
		Category:    "custom",
		Description: "Random friend name",
		Example:     "bill",
		Output:      "string",
		Generate: func(f *gofakeit.Faker, m *gofakeit.MapParams, info *gofakeit.Info) (any, error) {
			return f.RandomString([]string{"bill", "bob", "sally"}), nil
		},
	})
}
