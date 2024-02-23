package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := []rune(os.Args[0])
	for i := 2; i < len(args); i++ {
		z01.PrintRune(args[i])
	}
	z01.PrintRune('\n')
}
