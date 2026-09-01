package main

import "fmt"

func main() {
	var x [100]int
	// x é um array de 100 inteiros.
	// Como não atribuímos valores, todos começam com 0.

	fatia := x[0:10]
	// fatia é um slice que referencia as posições de 0 até 9 do array x.
	// Ele não cria uma cópia independente desses elementos.

	fmt.Printf("x = %T | fatia = %T\n", x, fatia)
	fmt.Println("Tamanho da fatia:", len(fatia))

	// Se alterarmos a fatia, podemos alterar o array x,
	// porque os dois estão usando os mesmos elementos nessa região.
	fatia[0] = 50

	fmt.Println("fatia:", fatia)
	fmt.Println("x[0]:", x[0])

	y := []int{}
	// y é um slice vazio.
	// Diferentemente de um array, seu tamanho pode mudar com append().

	fmt.Printf("y = %T\n", y)
	fmt.Println("Tamanho y:", len(y))

	for i := 1; i <= 200; i++ {
		y = append(y, i)
	}

	fmt.Println("Tamanho final de y:", len(y))
}
