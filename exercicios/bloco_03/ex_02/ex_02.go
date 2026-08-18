// faça o programa de uma sorveteria, onde o usuario por escolher:
// tipo de sorvete: casquinha (R$1,00), cascão (R$2,00), sundae (R$3,00)
// sabor do sorvete: chocolate, morango, creme, flocos
// cobertura: caramelo (R1,00), morango (R$1,50), chocolate (R$2,00), sem cobertura (R$0,00)
//Apresente o valor a ser pago

package main

import "fmt"

func main() {
	var tipoSorvete, saborSorvete, cobertura int
	fmt.Println("Escolha o tipo de sorvete:")
	fmt.Println("1 - Casquinha (R$ 1,00)")
	fmt.Println("2 - Cascão (R$ 2,00)")
	fmt.Println("3 - Sundae (R$ 3,00)")
	fmt.Scanf("%d", &tipoSorvete)

	fmt.Println("Escolha o sabor do sorvete:")
	fmt.Println("1 - Chocolate")
	fmt.Println("2 - Morango")
	fmt.Println("3 - Creme")
	fmt.Println("4 - Flocos")
	fmt.Scanf("%d", &saborSorvete)

	fmt.Println("Escolha a cobertura:")
	fmt.Println("1 - Caramelo (R$ 1,00)")
	fmt.Println("2 - Morango (R$ 1,50)")
	fmt.Println("3 - Chocolate (R$ 2,00)")
	fmt.Println("4 - Sem cobertura (R$ 0,00)")
	fmt.Scanf("%d", &cobertura)

	var precoTipoSorvete float64
	switch tipoSorvete {
	case 1:
		precoTipoSorvete = 1.00
	case 2:
		precoTipoSorvete = 2.00
	case 3:
		precoTipoSorvete = 3.00
	default:
		fmt.Println("Escolha inválida para o tipo de sorvete.")
		return
	}

	var precoCobertura float64
	switch cobertura {
	case 1:
		precoCobertura = 1.00
	case 2:
		precoCobertura = 1.50
	case 3:
		precoCobertura = 2.00
	case 4:
		precoCobertura = 0.00
	default:
		fmt.Println("Escolha inválida para a cobertura.")
		return
	}

	total := precoTipoSorvete + precoCobertura
	fmt.Printf("O valor total a ser pago é: R$ %.2f\n", total)
}
