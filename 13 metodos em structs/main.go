package main

import "fmt"

/* Definindo uma struct, com endereço como composição */
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

/* Na aula não se fez referência oa uso de '*' para ponteiros */
func (c *Cliente) desativar() {
	c.Ativo = false
}

/* Métodos em structs */
func main() {

	emerson := Cliente{
		Nome:  "Emerson",
		Idade: 41,
		Ativo: true,
	}

	emerson.desativar()

	fmt.Println(emerson)

}
