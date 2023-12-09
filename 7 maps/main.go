package main

import (
	"fmt"
)

func main() {

	salarios := map[string]int{"Emerson": 1000, "Mariana": 2000, "Marta": 3000}

	fmt.Println(salarios["Emerson"])

	// Apagando chave : valor
	delete(salarios, "Emerson")

	// Adicionando chave : valor
	salarios["Douglas"] = 5000

	fmt.Println(salarios)

	// Criando map vazio com make
	sal := make(map[string]int)

	// Outra forma de inicializar vazio
	sal1 := map[string]int{}

	sal["Emerson"] = 1000
	sal1["Marta"] = 100000

	// For sobre o map
	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	// For sobre o map com blank identifier
	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	}

}
