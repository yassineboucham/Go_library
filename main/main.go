package main

import (
	"os"
	"piscine"
)

func main() {
	args := os.Args[1:]

	arr := []string{}
	for _, arg := range args {
		arr = append(arr, piscine.r)
	}
}