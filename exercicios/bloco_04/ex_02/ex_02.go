// Faça um programa que receba 4 alturas usando um laço de repetição e realize a soma dessas alturas.

package main

import "fmt"

func main() {
	var altura float64
	fmt.Println("Entre com altura:")
	fmt.Scanf("%f", &altura)

	soma := 0.0
	for i := 0; i < 4; i++ {
		soma += altura
	}

	fmt.Printf("A soma das alturas é: %f\n", soma)
}

// %d é para inteiros, %f é para float, %s é para string, %v é para qualquer tipo de dado.
