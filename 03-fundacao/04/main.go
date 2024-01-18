package main

import "fmt"

type Endereco struct {
	Logradouro  string
	Numero      int
	Complemento string
	Cidade      string
	Estado      string
}

type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

// metodo em struct
func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("Cliente %s foi desativado!\n", c.Nome)
}

func DesativarPessoa(p Pessoa) {
	p.Desativar()
}

func main() {
	gabriel := Cliente{
		Nome:  "Gabriel",
		Idade: 21,
		Ativo: true,
	}
	gabriel.Logradouro = "Rua dos Bobos"
	fmt.Println(gabriel)
	DesativarPessoa(gabriel)
}
