package main

import "fmt"

func main() {

	var a int

	a = 10

	/* Mostra o endereço de memória através de '&' */
	fmt.Println(&a)

	/* * para se referir a ponteiro */
	var ponteiro *int = &a

	fmt.Println(ponteiro)

	/* Mudando o valor do ponteiro de 10 para 21 */
	*ponteiro = 21

	println(a)
}
