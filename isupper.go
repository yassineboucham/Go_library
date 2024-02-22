package piscine

func IsUpper(s string) bool {
	for _, run := range s {
		if run < 'A' || run > 'Z' {
			return false
		}
	}
	return true
}
