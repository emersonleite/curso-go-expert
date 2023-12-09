package main

import (
	"errors"
	"fmt"
)

func main() {
	f := sum(1, 2)
	fmt.Println(f)

	g, j := sum3(2, 50)
	fmt.Println(g, j)

	t, u := sumError(2, 50)
	fmt.Println(t, u)
}

func sum(a int, b int) int {
	return a + b
}

// Simplificando com uma declaração de tipo
func sum2(a, b int) int {
	return a + b
}

// Retornando mais de um valor
func sum3(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}
	return a + b, false
}

// Retornando mais de um valor com um deles sendo erro
func sumError(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	// nil -> equivalente a null
	return a + b, nil
}
