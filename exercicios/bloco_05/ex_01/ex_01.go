// faça um programa que receba 4 notas,
// calcule a media, minimo e maximo dessas notas e mostre na tela

package main

import (
	"fmt"
	"slices"
)

func main() {

	notas := [4]float64{}

	for i := 1; i <= 4; i++ {

		var nota float64
		fmt.Printf("Entre com a %da nota: ", i)
		fmt.Scanf("%f", &nota)
		notas[i-1] = nota
	}

	minimo := slices.Min(notas[:]) // com [:] eu consigo passar o array para a função Min, que recebe um slice
	maximo := slices.Max(notas[:]) // com [:] eu consigo passar o array para a função Max, que recebe um slice
	soma := 0.0
	for _, nota := range notas {
		soma += nota
	}
	media := soma / float64(len(notas))

	fmt.Printf("Média das notas: %.2f ", media)
	fmt.Printf("Nota mínima: %.2f ", minimo)
	fmt.Printf("Nota máxima: %.2f ", maximo)
}
