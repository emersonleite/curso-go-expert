package main

// Importando um pacote (padrão do Go)
import "fmt"

type ID int

var (
	f ID = 50
)

func main() {

	// %T traz o tipo na impressão

	fmt.Printf("O tipo de f é %T", f)

	// Pulando linha...
	fmt.Println()

	// %v traz o valor na impressão
	fmt.Printf("O valor de f é %v", f)
}
