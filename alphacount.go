package piscine

func AlphaCount(s string) int {
	nb := 0
	for _, run := range s {
		if run >= 'a' && run <= 'z' || run >= 'A' && run <= 'Z' {
			nb += 1
		}
	}
	return nb
}
