package main

import (
	"os"

	"github.com/01-edu/z01"
)

func repeatalpha(str string) {
	s := []rune(str)
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			for j := 0; j < int(s[i] - 96); j++ {
				z01.PrintRune(s[i])
			}
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			for j := 0; j < int(s[i] - 64); j++ {
				z01.PrintRune(s[i])
			}
		} else {
			z01.PrintRune(s[i])
		}
	}
	z01.PrintRune('\n')
}
 func main() {
	regs := os.Args[1]
	repeatalpha(regs)

 }