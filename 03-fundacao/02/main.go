package main

import "fmt"

func main() {
	salarios := map[string]int{
		"Gabriel": 1000,
		"João":    2000,
		"Maria":   3000,
	}
	fmt.Println(salarios["Gabriel"])
	delete(salarios, "Gabriel")
	fmt.Println(salarios)
	salarios["Pedro"] = 5000
	fmt.Println(salarios)

	for nome, salario := range salarios {
		fmt.Printf("%s recebe R$%d\n", nome, salario)
	}
}
