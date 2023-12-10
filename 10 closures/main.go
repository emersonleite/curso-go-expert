package main

import "fmt"

/* Closures */
func main() {

	func() {
		fmt.Println("Essa é uma função anônima")
	}()

	/* Utilizando closure (função dentro de função)  */
	total := func() int {
		return sum(1, 2, 3, 5, 36, 3, 5) * 2
	}()
	fmt.Println(total)

}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
