package main

import "fmt"

func main() {

	fmt.Println("Entre com um valor:")
	var x float64
	fmt.Scanf("%f", &x) // eu tenho que passar o endereço da variavel para a função Scanf, por isso o uso do & antes da variavel x

	fmt.Println("Você escolheu o valor", x)
}
