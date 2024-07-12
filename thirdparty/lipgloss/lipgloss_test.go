package lipgloss

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
	"github.com/charmbracelet/lipgloss/table"
)

func TestStyle(t *testing.T) {
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingTop(2).
		PaddingLeft(4).
		Width(22)

	t.Log("\n" + style.Render("Hello, lipgloss!"))
}

func TestColor(t *testing.T) {
	f := func(fgColor, message string) {
		style := lipgloss.NewStyle().Foreground(lipgloss.Color(fgColor))
		outputText := strings.Repeat(message+" ", 5)
		t.Log(style.Render(outputText))
	}

	// ANSI 16 colors (4-bit)
	f("5", "magenta")
	f("9", "red")
	f("12", "lightBlue")

	// ANSI 256 colors (8-bit)
	f("86", "aqua")
	f("201", "hotPink")
	f("202", "orange")

	// True Color (16,777,216 colors; 24-bit)
	f("#0000FF", "blue")
	f("#04B575", "green")
	f("#3C3C3C", "darkGrey")
}

func TestANSIFormatting(t *testing.T) {
	style := lipgloss.NewStyle().
		Bold(true).
		Italic(true).
		Faint(true).
		Blink(true).
		Strikethrough(true).
		Underline(true).
		Reverse(true)

	t.Log(style.Render("Lipgloss Formatting"))
}

func TestWidthAndHeight(t *testing.T) {
	style := lipgloss.NewStyle().
		SetString("What’s for lunch?").
		Width(24).
		Height(5).
		Foreground(lipgloss.Color("63"))

	t.Log(style)
}

func TestTable(t *testing.T) {
	headerStyle := lipgloss.NewStyle().Bold(true)
	evenRowStyle := lipgloss.NewStyle().Background(lipgloss.Color("#555555"))
	oddRowStyle := lipgloss.NewStyle()

	rows := [][]string{
		{"Chinese", "您好", "你好"},
		{"Japanese", "こんにちは", "やあ"},
		{"Arabic", "أهلين", "أهلا"},
		{"Russian", "Здравствуйте", "Привет"},
		{"Spanish", "Hola", "¿Qué tal?"},
	}

	tb := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == 0:
				return headerStyle
			case row%2 == 0:
				return evenRowStyle
			default:
				return oddRowStyle
			}
		}).
		Headers("LANGUAGE", "FORMAL", "INFORMAL").
		Rows(rows...)

	// You can also add tables row-by-row
	tb.Row("English", "You look absolutely fabulous.", "How's it going?")

	t.Logf("\n%v", tb)
}

func TestList(t *testing.T) {
	got := list.New(
		"root",
		list.New(
			"child1",
			"child2",
			list.New(
				"child 3",
				list.New(
					"child 3.1",
				),
			),
		),
	)

	t.Log(got)
}
