// faça um exercicio que receba um numero inteiro e calcule a raiz quadrada e exiba o resultado

package main

import (
	"fmt"
	"math"
)

func main() {
	var numero int
	fmt.Println("Digite um número inteiro:")
	fmt.Scanf("%d", &numero)

	raiz := math.Sqrt(float64(numero))
	fmt.Printf("A raiz quadrada de %d é %.2f\n", numero, raiz)
	// o \n serve para pular uma linha no final da saída
}
