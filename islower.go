package piscine

func IsLower(s string) bool {
	for _, run := range s {
		if run < 'a' || run > 'z' {
			return false
		}
	}
	return true
}
