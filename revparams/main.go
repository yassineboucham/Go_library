package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for i := len(args) - 1; i > 0; i-- {
		for j := 0; j < len(args[i]); j++ {
			z01.PrintRune(rune(args[i][j]))
		}
		z01.PrintRune('\n')
	}
}
