package piscine

func ToUpper(s string) string {
	str := []rune(s)
	for i, run := range s {
		if run >= 'a' && run <= 'z' {
			str[i] -= 32
		} else {
			continue
		}
	}
	return string(str)
}
