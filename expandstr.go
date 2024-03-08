package piscine

import (
	"fmt"
	"os"
)

func expandstr() {
	args := os.Args
	str := ""
	if len(args) != 2 {
		return
	}
	for j := 0; j < len(args[1]); j++ {
		if string(args[1][j]) != " " {
			str += string(args[1][j])
		} else {
			x := 0
			for i := j + 1; i < len(args[1]); i++ {
				if args[1][i] == 32 {
					x++
					continue
				} else {
					str += "."
					j += x
					break
				}
			}
		}
	}
	arr := ""
	for i := 0; i < len(str); i++ {
		if str[i] == '.' {
			arr += "   "
		} else {
			arr += string(str[i])
		}
	}
	fmt.Println(arr)
}
