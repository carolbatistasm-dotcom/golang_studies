package main

import "fmt"

func main() {
	var dia int
	fmt.Println("Entre com o número do dia (1-7):")
	fmt.Scanf("%d", &dia)

	switch dia {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda-feira")
	case 3:
		fmt.Println("Terça-feira")
	case 4:
		fmt.Println("Quarta-feira")
	case 5:
		fmt.Println("Quinta-feira")
	case 6:
		fmt.Println("Sexta-feira")
	case 7:
		fmt.Println("Sábado")
	default:
		fmt.Println("Número inválido. Por favor, entre com um número entre 1 e 7.")
	}
}
