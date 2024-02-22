package piscine

func StrRev(s string) string {
	str := []rune(s)
	j := len(str) - 1
	for i := 0; i < j; i++ {
		t := str[i]
		str[i] = str[j]
		str[j] = t
		j--
	}
	return string(str)
}
