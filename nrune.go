package piscine

func NRune(s string, n int) rune {
	r := []rune(s)
	if n > len(s) || n <= 0 {
		return 0
	} else {
		return r[n-1]
	}
}
