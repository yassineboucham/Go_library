package piscine

import "fmt"

func FindNextPrime(nb int) int {
	for i := 2; i < 999900; i++ {
		if nb%i != 0 && i < nb {
			return nb
		} else if i > nb {
			for j := nb; j < 999900; j++ {
				if i%j != 0 && i < j {
					return i
				} else if i > j-1 {
					continue
				}
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(FindNextPrime(17))
	fmt.Println(FindNextPrime(4))
}
