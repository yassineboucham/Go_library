package main

import (
	"fmt"
	"os"
)

func UppToLow(str string) string {
	rev := ""
	for _, s := range str {
		if s >= 'a' && s <= 'z' {
			rev += string(s - 32)
		} else if s >= 'A' && s <= 'Z' {
			rev += string(s + 32)
		} else {
			rev += string(s)
		}
	}
	return rev
}

func main() {
	args := os.Args[1:]
	arr := []string{}
	for _, arg := range args {
		arr = append(arr, UppToLow(arg))
	}
	for _, str := range arr {
		fmt.Println(str)
	}
}
