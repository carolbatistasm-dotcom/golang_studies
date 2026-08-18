// faça um programa que calcule a média de 3 notas
// e verifique se o aluno passou de ano
// a nota de corte será 6

package main

import "fmt"

func main() {

	var nota1, nota2, nota3 float64
	fmt.Println("Digite a primeira nota:")
	fmt.Scanf("%f", &nota1)
	fmt.Println("Digite a segunda nota:")
	fmt.Scanf("%f", &nota2)
	fmt.Println("Digite a terceira nota:")
	fmt.Scanf("%f", &nota3)

	media := (nota1 + nota2 + nota3) / 3.0

	if media >= 6 {
		fmt.Println("Parabéns! Você passou de ano com média:", media)
	} else {
		fmt.Println("Infelizmente você não passou de ano. Sua média foi:", media)
	}
}
