package main

type MyNumber int

// constrante para tipo genérico
// ~int, o ~ significa que o go vai aceitar qualquer tipo que seja int
type Number interface {
	~int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Wesley": 10, "Ana": 20, "João": 30}
	m2 := map[string]float64{"Wesley": 10.0, "Ana": 20.0, "João": 30.0}
	m3 := map[string]MyNumber{"Wesley": 10, "Ana": 20, "João": 30}
	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10.0))
}
