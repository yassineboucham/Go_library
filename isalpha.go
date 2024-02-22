package piscine

func IsAlpha(s string) bool {
	for _, run := range s {
		if run >= 'a' && run <= 'z' || run >= 'A' && run <= 'Z' || run >= '0' && run <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
