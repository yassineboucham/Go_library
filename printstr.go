package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	str := []rune(s)
	for index := range str {
		z01.PrintRune(str[index])
	}
}
