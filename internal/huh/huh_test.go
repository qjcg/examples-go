//go:build e2e

package huh

import (
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
