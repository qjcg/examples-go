package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/charmbracelet/lipgloss"
)

const (
	ansiMagenta = "5"
	ansiCyan    = "6"
)

var (
	cyan    = lipgloss.NewStyle().Foreground(lipgloss.Color(ansiCyan))
	magenta = lipgloss.NewStyle().Foreground(lipgloss.Color(ansiMagenta))
)

// TODO: Decide whether to use this.
// type Counter struct {
// 	w      io.Writer // TODO: Consider chan(int) instead?
// 	Start  int
// 	End    int
// 	Symbol string
// 	Style  lipgloss.Style
// }

// func NewCounter() *Counter {
// 	return &Counter{}
// }

// func (c *Counter) Count() {
// 	switch {
// 	case c.Start > c.End:
// 		for i := c.Start; i >= c.End; i-- {
// 		}
// 	case c.Start < c.End:
// 		for i := c.Start; i <= c.End; i++ {
// 		}
// 	case c.Start == c.End:
// 		return
// 	}
// }

// Up prints all numbers from 0-100, counting Up.
func Up(w io.Writer) {
	for i := 0; i <= 100; i++ {
		fmt.Fprintf(w, cyan.Render("U%d "), i)
	}
}

// Down prints all numbers from 100-0, counting Down.
func Down(w io.Writer) {
	for i := 100; i >= 0; i-- {
		fmt.Fprintf(w, magenta.Render("D%d "), i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		Up(os.Stdout)
		wg.Done()
	}()

	go func() {
		Down(os.Stdout)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println()
}
