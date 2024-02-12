package main

import "fmt"

type Cliente struct {
	nome string
}

func (c Cliente) andou() {
	c.nome = "Emerson Leite"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

func (c *Cliente) andouComPonteiro() {
	c.nome = "Emerson Leite"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

func main() {

	emerson := Cliente{
		nome: "Emerson",
	}

	/* Aqui a passagem de parâmetro é por cópia */
	emerson.andou()
	fmt.Printf("O valor da struct com nome é %v \n\n", emerson.nome)

	/* Aqui a função 'andouComPonteiro' muda o próprio valor da referência de nome na memória */
	emerson.andouComPonteiro()
	fmt.Printf("O valor da struct com nome é %v \n", emerson.nome)

}
