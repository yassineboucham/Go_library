package piscine

func RecursiveFactorial(nb int) int {
	if nb == 0 {
		return 1
	}
	if nb <= 20 && nb >= 0 {
		return nb * RecursiveFactorial(nb-1)
	}
	return 0
}