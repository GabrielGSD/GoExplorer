package main

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	/*
		// Ponteiro é utilizado quando não quero que o que estou criando esteja em um escopo local
		// Memória -> Endereço -> Valor
		a := 10
		var ponteiro *int = &a
		*ponteiro = 20
		b := &a
		*b = 30
		println(a, *b)
	*/

	minhaVar1 := 10
	minhaVar2 := 20
	println(soma(&minhaVar1, &minhaVar2))
	println(minhaVar1, minhaVar2)
}
