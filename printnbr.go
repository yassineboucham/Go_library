package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	arr := []rune{}
	var xx string = "-9223372036854775808"
	if n == -9223372036854775808 {
		for i := 0; i < 20; i++ {
			z01.PrintRune(rune(xx[i]))
		}
		return
	} else if n < 0 {
		z01.PrintRune('-')
		n = n * -1
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	for n != 0 {
		arr = append(arr, rune((n%10)+'0'))
		n = n / 10
	}
	for i := len(arr) - 1; i >= 0; i-- {
		z01.PrintRune(arr[i])
	}
}
