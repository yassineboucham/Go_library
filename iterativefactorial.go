package piscine

func IterativeFactorial(nb int) int {
	res := 1
	if nb > 20 || nb < 0 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		res *= i
	}
	return res
}
