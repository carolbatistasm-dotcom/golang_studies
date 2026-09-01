package main

import "fmt"

func main() {

	var idade int
	fmt.Println("Entre com a sua idade:")
	fmt.Scanf("%d", &idade) // "%d" é usado para ler um número inteiro

	if idade >= 60 {
		fmt.Println("Você é idoso.")
	} else if idade >= 18 { // elif no go é explícito com else if
		fmt.Println("Você é maior de idade.")
	} else {
		fmt.Println("Você é menor de idade.")
	}
}
