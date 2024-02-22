package piscine

func IsNumeric(s string) bool {
	for _, run := range s {
		if run < '0' || run > '9' {
			return false
		}
	}
	return true
}
