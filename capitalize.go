package piscine

func isalphanum(run rune) bool {
	if run >= 'a' && run <= 'z' || run >= 'A' && run <= 'Z' || run >= '0' && run <= '9' {
		return true
	} else {
		return false
	}
}

func isalpha(run rune) bool {
	if run >= 'a' && run <= 'z' || run >= 'A' && run <= 'Z' {
		return true
	} else {
		return false
	}
}

func toupper(run rune) rune {
	if run >= 'A' && run <= 'Z' {
		return run
	} else {
		return run - 32
	}
}

func Capitalize(s string) string {
	str := []rune(s)
	str[0] = toupper(str[0])
	for i := 1; i < len(s); i++ {
		if isalpha(str[i]) {
			if !isalphanum(str[i - 1]){
				str[i] = toupper(str[i])
			}
		}
	}
	return string(str)
}
