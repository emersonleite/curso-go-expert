package main

import "fmt"

func main() {

	//var a int

	a := 10

	/* Mostra o endereço de memória através de '&' */
	fmt.Println(&a)

	/* asterisco '*' para se referir a ponteiro */
	var ponteiro *int = &a

	// Mostrando endereço na memória de ponteiro
	fmt.Println(ponteiro)

	// Outra forma de criar um ponteiro sem definir o tipo explicitamente
	ponteiro2 := &a

	// Mostrando endereço na memória de ponteiro2
	fmt.Println(ponteiro2)

	/* Mudando o valor do ponteiro de 10 para 21 e depois, através o ponteiro2 para 22 */
	*ponteiro = 21

	*ponteiro2 = 22

	println("valor de a = ", a)

	/* Para mostrar o valor do ponteiro utiliza-se o asterisco '*' */
	fmt.Println("valor de ponteiro = ", *ponteiro)
}
