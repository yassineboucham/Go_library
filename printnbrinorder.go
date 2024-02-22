package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	arr := []int{}
	swap := 0
	if n < 9 {
		z01.PrintRune(rune(n + '0'))
	}
	for n > 0 {
		nb := n % 10
		n = n / 10
		arr = append(arr, nb)
	}
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				swap = arr[i]
				arr[i] = arr[j]
				arr[j] = swap
			}
		}
		z01.PrintRune(rune(arr[i] + '0'))
	}
}
