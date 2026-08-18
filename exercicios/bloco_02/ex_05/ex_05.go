// faça um programa que exiba o dobro de um numero insirido pelo usuario

package main

import "fmt"

func main() {
	var x float64

	fmt.Println("Entre com um número:")
	fmt.Scanf("%f", &x)

	resultado := x * 2
	fmt.Printf("O dobro de %.2f é %.2f\n", x, resultado)

	// toda vez que eu quero que o usuario entre com um valor, uso println e scanf
	// scanf é usado para ler o valor digitado pelo usuario e armazenar na variavel x
}
