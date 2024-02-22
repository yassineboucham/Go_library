package piscine

func Swap(a *int, b *int) {
	nb := *a
	*a = *b
	*b = nb
}
