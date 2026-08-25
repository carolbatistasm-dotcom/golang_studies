// faça um programa que receba 6 temperaturas
// remova a primeira e a ultima para calcular a media.
// se a media for acima de 30 graaus, exiba que houve um aumento da temperatura

package main

import (
	"fmt"
)

func main() {

	temperaturas := [6]float64{}

	for i := 1; i <= 6; i++ {

		var temperatura float64
		fmt.Printf("Entre com a %da temperatura: ", i)
		fmt.Scanf("%f", &temperatura)
		temperaturas[i-1] = temperatura
	}

	// removendo a primeira e a ultima temperatura
	nova_temp := temperaturas[1 : len(temperaturas)-1]

	soma := 0.0
	for _, temperatura := range nova_temp {
		soma += temperatura
	}
	media := soma / float64(len(nova_temp))

	fmt.Printf("Média das temperaturas: %.2f ", media)

	if media > 30 {
		fmt.Println("Houve um aumento da temperatura.")
	} else {
		fmt.Println("Não houve aumento da temperatura.")
	}
}
