package main

import "fmt"

/* Funções variádicas */
func main() {

	t := sum(1, 2, 3, 4, 5, 6, 7)

	fmt.Println(t)

}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
