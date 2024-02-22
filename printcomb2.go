package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for nb1 := '0'; nb1 <= '9'; nb1++ {
		for nb2 := '0'; nb2 <= '9'; nb2++ {
			for nb3 := '0'; nb3 <= '9'; nb3++ {
				for nb4 := '0'; nb4 <= '9'; nb4++ {
					if ((nb1-'0')*10 + nb2) < ((nb3-'0')*10 + nb4) {
						z01.PrintRune(nb1)
						z01.PrintRune(nb2)
						z01.PrintRune(' ')
						z01.PrintRune(nb3)
						z01.PrintRune(nb4)
						if nb1 == '9' && nb2 == '8' && nb3 == '9' && nb4 == '9' {
							break
						}
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
