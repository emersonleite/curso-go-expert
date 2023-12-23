package main

import (
	"fmt"
)

/* Definindo uma struct */
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

/* Structs */
func main() {

	emerson := Cliente{
		Nome:  "Emerson",
		Idade: 41,
		Ativo: true,
	}

	fmt.Printf("Nome %s, Idade: %d, Ativo: %t", emerson.Nome, emerson.Idade, emerson.Ativo)

	// Mudando valor de uma struct
	emerson.Ativo = false

	fmt.Println()
	fmt.Println(emerson.Ativo)

}
