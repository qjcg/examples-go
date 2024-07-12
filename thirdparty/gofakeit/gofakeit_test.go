package gofakeit

import (
	"fmt"
	"log"

	"github.com/brianvoe/gofakeit/v7"
)

func init() {
	// Using a constant seed for all examples makes their outputs deterministic.
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
	// 42
	// Frances
	// terrenceheathcote@marquardt.org
}

// Custom function via AddFuncLookup.
func ExampleAddFuncLookup() {
	gofakeit.AddFuncLookup("chilli", gofakeit.Info{
		Category:    "custom",
		Description: "Random chilli pepper name",
		Example:     "habanero",
		Output:      "string",
		Generate: func(f *gofakeit.Faker, m *gofakeit.MapParams, info *gofakeit.Info) (any, error) {
			return f.RandomString([]string{"tabasco", "habanero", "scotch bonnet"}), nil
		},
	})

	type Foo struct {
		Pepper string `fake:"{chilli}"`
	}

	var f Foo
	err := gofakeit.Struct(&f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(f.Pepper)

	// Output: tabasco
}

// Custom function taking param(s) via AddFuncLookup.
func ExampleAddFuncLookup_withParams() {
	gofakeit.AddFuncLookup("dominos", gofakeit.Info{
		Category:    "",
		Description: "Random dominos.",
		Example:     "",
		Output:      "string",
		Params: []gofakeit.Param{
			{Field: "count", Type: "uint", Default: "5", Description: "Number of dominos to generate"},
		},
		Generate: func(f *gofakeit.Faker, m *gofakeit.MapParams, info *gofakeit.Info) (any, error) {
			count, err := info.GetUint(m, "count")
			if err != nil {
				return nil, err
			}

			var dominos string
			start := 0x1f030
			end := 0x1f09f
			for i := 0; i < int(count); i++ {
				v := gofakeit.Number(start, end) // Uses random seed from init().
				dominos += string(rune(v))
			}

			return dominos, nil
		},
	})

	type Foo struct {
		Dominos          string `fake:"{dominos}"`
		DominosWithCount string `fake:"{dominos:10}"`
	}

	var f Foo
	err := gofakeit.Struct(&f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(f.Dominos)
	fmt.Println(f.DominosWithCount)

	// Output:
	// 🂈🁝🂉🀻🂌
	// 🁇🀵🁎🁷🂂🀺🂌🁍🁞🂉
}
