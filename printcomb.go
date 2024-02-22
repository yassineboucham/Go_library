package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for n1 := '0'; n1 <= '7'; n1++ {
		for n2 := n1 + 1; n2 <= '8'; n2++ {
			for n3 := n2 + 1; n3 <= '9'; n3++ {
				z01.PrintRune(n1)
				z01.PrintRune(n2)
				z01.PrintRune(n3)
				if n1 != '7' || n2 != '8' || n3 != '9' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
