package main

import "fmt"

/* Definindo uma struct, com endereço como composição */
type ClienteTipo1 struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

/* Também pode ser definido assim, com address sendo do tipo endereço.
 */
type ClienteTipo2 struct {
	Nome  string
	Idade int
	Ativo bool
	End   Endereco
}

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Estado string
}

/* Composição de structs */
func main() {

	endereco := Endereco{
		Rua:    "Fernandes de Barros, 1732",
		Numero: 1732,
		Cidade: "Curitóba",
		Estado: "PR",
	}

	emerson := ClienteTipo1{
		Nome:     "Emerson",
		Idade:    41,
		Ativo:    true,
		Endereco: endereco,
	}

	mariana := ClienteTipo2{
		Nome:  "Mariana",
		Idade: 18,
		Ativo: true,
		End:   endereco,
	}

	/* Mudança de endereço */
	emerson.Endereco.Cidade = "Curitóba"

	fmt.Println(emerson)
	fmt.Println(endereco)
	fmt.Println(mariana)

}
