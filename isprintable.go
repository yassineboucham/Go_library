package piscine

func IsPrintable(s string) bool {
	for _, run := range s {
		if run < 32 || run > 127 {
			return false
		}
	}
	return true
}
