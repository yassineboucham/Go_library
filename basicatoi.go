package piscine

func BasicAtoi(s string) int {
	l := len(s)
	num := 0
	for i := 0; i < l; i++ {
		if s[i] < '9' || s[i] > '0' {
			if i == l-1 {
				num = num + (int(s[i]) - '0')
			} else {
				num = (num + int(s[i]) - '0') * 10
			}
		} else {
			num = 0
		}
	}
	return num
}
