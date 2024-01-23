package main

type Conta struct {
	saldo int
}

// factory function
func NewConta() *Conta {
	return &Conta{saldo: 0}
}

// Quando coloco * eu não estou passando uma cópia, estou passando o endereço de memória (a alteração reflete no objeto original)
func (c *Conta) simular(valor int) int {
	c.saldo += valor
	return c.saldo
}

func main() {
	conta := NewConta()
	conta.saldo = 100
	println(conta.saldo)
	conta.simular(50)
	println(conta.saldo)
}
