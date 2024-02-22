package piscine

func UltimateDivMod(a *int, b *int) {
	nb1 := *a
	nb2 := *b
	*a = nb1 / nb2
	*b = nb1 % nb2
}
