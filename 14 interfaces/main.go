package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

type Empresa struct {
	Nome     string
	Fundacao int
	Ativo    bool
}

type Pessoa interface {
	Desativar()
}

func (cliente *Cliente) Desativar() {
	cliente.Ativo = false
}

func (empresa *Empresa) Desativar() {
	empresa.Ativo = false
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {

	emerson := Cliente{
		Nome:  "Emerson",
		Idade: 42,
		Ativo: true,
	}

	tabajara := Empresa{
		Nome:     "Tabajara",
		Fundacao: 52,
		Ativo:    true,
	}

	Desativacao(&emerson)
	Desativacao(&tabajara)

	fmt.Println(emerson)
	fmt.Println(tabajara)

}
