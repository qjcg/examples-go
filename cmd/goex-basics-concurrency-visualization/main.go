package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/fatih/color"
)

var (
	cyan    = color.New(color.FgCyan)
	magenta = color.New(color.FgMagenta)
)

// Up prints all numbers from 0-100, counting Up.
func Up(w io.Writer) {
	for i := 0; i <= 100; i++ {
		cyan.Fprintf(w, "U%d ", i)
	}
}

// Down prints all numbers from 100-0, counting Down.
func Down(w io.Writer) {
	for i := 100; i >= 0; i-- {
		magenta.Fprintf(w, "D%d ", i)
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
