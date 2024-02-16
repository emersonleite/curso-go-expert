package main

import "fmt"

/* Slices */
func main() {
	// slice sem tamanho fixo

	s := []int{10, 20, 30, 50, 60, 70, 80, 90, 100}

	// cap é a capacidade
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// Apagando itens
	// Saída --->  // len=0 cap=9 []
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	// Deixando 4 itens, e utilizando os primeiros itens
	// Saída ---> len=4 cap=9 [10 20 30 50]
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	// Ignorando itens do começo, eliminando-se os 2 primeiros itens
	// A capacidade é diminuida
	// Saída ---> len=7 cap=7 [30 50 60 70 80 90 100]
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	// Aumentando a capacidade ao se adicionar um item
	// Saída ---> len=10 cap=20 [10 20 30 50 60 70 80 90 100 110]
	s = append(s, 110)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}
