package main

import (
	"fmt"
)

func soma(a, b int) int {
	a = 50
	return a + b
}

// Passando o endereço na memória
func somaComPonteiro(a, b *int) int {
	// Mudando o valor da referência na memória
	*a = 50
	return *a + *b
}

func main() {

	minhaVar1 := 10
	minhaVar2 := 20

	somaDoValor := soma(minhaVar1, minhaVar2)
	fmt.Println(somaDoValor)
	// minhaVar1 continua sendo 10, pois a função não mudou o valor da referência de memória
	fmt.Println(minhaVar1)

	// Passando a referência com '&'
	somaValorComPonteiro := somaComPonteiro(&minhaVar1, &minhaVar2)
	fmt.Println(somaValorComPonteiro)
	// minhaVar1 passa a ser 50, pois a função somaComPonteiro a muda
	fmt.Println(minhaVar1)
}
