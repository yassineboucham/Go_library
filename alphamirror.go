package main

import (
	"fmt"
	"os"
)

func alphamirror(str string) string {
	s := []rune(str)
	for i := 0; i < len(str); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			s[i] = 219 - s[i]
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			s[i] = 'A'+'Z' - s[i]
		}
	}
	return string(s)
}

func main() {
	regs := os.Args
	if len(regs) == 2 {
		fmt.Println(alphamirror(regs[1]))
	}
}
