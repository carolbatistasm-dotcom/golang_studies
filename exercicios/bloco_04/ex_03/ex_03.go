//Faça um programa que receba uma quantidade indefinida de valores correspondentes a “saldo em conta”,
// mas quando o usuário apertar “enter” sem digitar valor algum, o programa para de receber valores,
// e exibe a soma de todos os valores digitados anteriormente
// sem usar breakio

package main

import (
	"fmt"
	"strconv"
)

func main() {

	var soma float64

	for {

		var entrada string

		fmt.Print("Entre com o valor:")
		fmt.Scanf("%s", &entrada)

		fmt.Println(entrada)

		if entrada == "" {
			break
		}

		valor, error := strconv.ParseFloat(entrada, 64)

		if error != nil { // nill é um valor vazio
			fmt.Println("Entre com um valor válido!")
			continue
		}

		soma += valor
	}

	fmt.Printf("Saldo total em conta: R$%.2f\n", soma)

}
