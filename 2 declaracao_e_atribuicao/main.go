package main

// Declarando constante
const a = 10

// Aqui fora de main o escopo é global, acessível de qualquer função
// Declarando variáveis
var (
	b bool
	c int
	d string
	e float64
	// declaração e atribuição
	f bool = true
)

func main() {

	// Escopo local da variável
	// Variáveis de escopo local declaradas sempre devem ser utilizadas
	var a string = "Oh loco"

	// Declaração, atribuição e inferência. Na primeira vez se utiliza o :. Para reatribuir depois, não precisa do : .
	g := "X"

	println(g)

	println(a)

	println("Ola mosca")

	// Se não se associar valor, o valor inferido nesse caso é false
	println(b)

	// Se não se associar valor, o valor inferido nesse caso é 0
	println(c)

	// valor em branco
	println(d)

	// +0.000000e+000
	println(e)

	println(f)

}
