package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[0]
	fmt.Printf("%s\n", args[2:])
}
