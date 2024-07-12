//go:build interactive

package huh

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/charmbracelet/huh"
)

func TestNewInput(t *testing.T) {
	var name string

	huh.NewInput().
		Title("What's your name?").
		Value(&name).
		Run()

	t.Logf("Hello, %s!\n", name)
}

func TestNewSelect(t *testing.T) {
	var country string

	huh.NewSelect[string]().
		Title("Pick a country.").
		Options(
			huh.NewOption("United States", "US"),
			huh.NewOption("Germany", "DE"),
			huh.NewOption("Brazil", "BR"),
			huh.NewOption("Canada", "CA"),
		).
		Value(&country).
		Run()

	t.Logf("Country selected: %s\n", country)
}

func TestNewMultiSelect(t *testing.T) {
	var toppings []string

	huh.NewMultiSelect[string]().
		Options(
			huh.NewOption("Lettuce", "Lettuce").Selected(true),
			huh.NewOption("Tomatoes", "Tomatoes").Selected(true),
			huh.NewOption("Charm Sauce", "Charm Sauce"),
			huh.NewOption("Jalapeños", "Jalapeños"),
			huh.NewOption("Cheese", "Cheese"),
			huh.NewOption("Vegan Cheese", "Vegan Cheese"),
			huh.NewOption("Nutella", "Nutella"),
		).
		Title("Toppings").
		Limit(4).
		Value(&toppings).
		Run()

	t.Logf("Toppings selected: %s\n", strings.Join(toppings, ", "))
}

func TestNewConfirm(t *testing.T) {
	var confirm bool

	huh.NewConfirm().
		Title("Are you sure?").
		Affirmative("Yes!").
		Negative("No.").
		Value(&confirm).
		Run()

	t.Logf("Are you sure? %v\n", confirm)
}

func TestNewForm(t *testing.T) {
	type Order struct {
		burger       string
		name         string
		instructions string
		toppings     []string
		discount     bool
	}

	String := func(o *Order) string {
		return fmt.Sprintf("\nBurger: %s\nName: %s\nInstructions: %s\nToppings: %s\nDiscount: %v\n", o.burger, o.name, o.instructions, o.toppings, o.discount)
	}

	var order Order

	// TODO: ensure input is a valid name.
	validateName := func(s string) error { return nil }

	form := huh.NewForm(
		// Prompt the user to choose a burger.
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(
					huh.NewOption("Charmburger Classic", "classic"),
					huh.NewOption("Chickwich", "chickwich"),
					huh.NewOption("Fishburger", "Fishburger"),
					huh.NewOption("Charmpossible™ Burger", "charmpossible"),
				).
				Title("Choose your burger").
				Value(&order.burger),
		),

		// Prompt for toppings and special instructions.
		// The customer can ask for up to 4 toppings.
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Options(
					huh.NewOption("Lettuce", "Lettuce").Selected(true),
					huh.NewOption("Tomatoes", "Tomatoes").Selected(true),
					huh.NewOption("Charm Sauce", "Charm Sauce"),
					huh.NewOption("Jalapeños", "Jalapeños"),
					huh.NewOption("Cheese", "Cheese"),
					huh.NewOption("Vegan Cheese", "Vegan Cheese"),
					huh.NewOption("Nutella", "Nutella"),
				).
				Title("Toppings").
				Limit(4).
				Value(&order.toppings),
		),

		// Gather final details for the order.
		huh.NewGroup(
			huh.NewInput().
				Title("What's your name?").
				Value(&order.name).
				Validate(validateName),

			huh.NewText().
				Title("Special Instructions").
				Value(&order.instructions).
				CharLimit(400),

			huh.NewConfirm().
				Title("Would you like 15% off").
				Value(&order.discount),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	t.Log(String(&order))
}
