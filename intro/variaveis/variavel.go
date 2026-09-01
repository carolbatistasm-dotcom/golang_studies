package main

import "fmt"

// para criar variaveis tem que declarar o tipo da variavel, e para declarar uma variavel tem que usar a palavra var
// exemplo: var nomeDaVariavel tipoDaVariavel
// no go a gente usa pascal case para nomear variaveis, ou seja,
// a primeira letra de cada palavra é maiuscula e não se usa underline

func main() {

	var nome string = "Gustavo"
	var idade int = 20
	var altura float32 = 1.75
	var peso float64 = 70.5
	var ativo bool = true

	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Altura:", altura)
	fmt.Println("Peso:", peso)
	fmt.Println("Ativo:", ativo)

	// eu posso só declarar a variavel sem atribuir um valor para ela
	var x int
	fmt.Println("Valor de x:", x) // quando eu nao atribuo um valor para a variavel, o go atribui o valor zero para a variavel

	var z, w, y float64
	fmt.Printf("z = %f | w = %f | y = %f \n", z, w, y)

	// a diferença entre o println e o printf é que o println adiciona uma quebra de linha no final da mensagem,
	// enquanto o printf não adiciona a quebra de linha no final da mensagem
	// e o printf permite formatar a mensagem, enquanto o println não permite formatar a mensagem

	var a, b, c string = "Carol", "Batista", "Simões"
	fmt.Printf("a = %s | b = %s | c = %s \n", a, b, c)
	fmt.Println(a, b, c)

	// declaracao curta de variaveis, onde o tipo da variavel é inferido pelo valor atribuido a ela
	d := 10 // eu posso até reatribuir um valor para a variavel, mas não posso mudar o tipo da variavel
	e := 20.5
	f := "Gustavo"
	g := true

	fmt.Println("d:", d)
	fmt.Println("e:", e)
	fmt.Println("f:", f)
	fmt.Println("g:", g)

	fmt.Printf("d = %T ; e = %T ; f = %T ; g = %T \n", d, e, f, g)
}
