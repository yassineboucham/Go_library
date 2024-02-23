package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			z01.PrintRune(rune(arr[i][j]))
		}
		z01.PrintRune('\n')
	}
}
