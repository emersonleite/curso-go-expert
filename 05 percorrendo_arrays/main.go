package main

import "fmt"

var (
	b bool
	c int
	d string  = "Emerson"
	e float64 = 1.2
	// declaração e atribuição
	f bool = true
)

/* Percorrendo arrays  */
func main() {

	// Definindo array com posições fixas
	var meuArray [3]int

	//  Adicionando elementos no array
	meuArray[0] = 1
	meuArray[2] = 2
	meuArray[1] = 1

	fmt.Println(len(meuArray) - 1)

	// Percorrendo  array
	for i, v := range meuArray {
		// %d representar valor de int
		fmt.Printf("O valor do índice %d é %d\n", i, v)
	}

}
