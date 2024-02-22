package piscine

func StrLen(s string) int {
	tab := []rune(s)
	j := 0
	for i := range tab {
		j = i
	}
	return j + 1
}
