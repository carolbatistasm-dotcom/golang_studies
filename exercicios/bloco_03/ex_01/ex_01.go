// faça um programa que vende uma garrafa de água:
// se o cliente escolher agua mineral natural, o preço é R$ 2,00
// se o cliente escolher agua mineral com gás, o preço é R$ 2,50
// se o cliente escolher agua de coco, o preço é R$ 3,00
// se o cliente escolher agua saborizada, o preço é R$ 4,00

package main

import "fmt"

func main() {
	var escolha, quantidade int
	fmt.Println("Escolha o tipo de água:")
	fmt.Println("1 - Água mineral natural (R$ 2,00)")
	fmt.Println("2 - Água mineral com gás (R$ 2,50)")
	fmt.Println("3 - Água de coco (R$ 3,00)")
	fmt.Println("4 - Água saborizada (R$ 4,00)")
	fmt.Scanf("%d", &escolha)

	fmt.Println("Digite a quantidade de garrafas que deseja comprar:")
	fmt.Scanf("%d", &quantidade)

	txtTemplate := "O valor ficou: R$ %.2f x %d = R$%.2f\n"

	switch escolha {
	case 1:
		fmt.Printf(txtTemplate, 2.00, quantidade, 2.00*float64(quantidade))
	case 2:
		fmt.Printf(txtTemplate, 2.50, quantidade, 2.50*float64(quantidade))
	case 3:
		fmt.Printf(txtTemplate, 3.00, quantidade, 3.00*float64(quantidade))
	case 4:
		fmt.Printf(txtTemplate, 4.00, quantidade, 4.00*float64(quantidade))
	default:
		fmt.Println("Escolha inválida. Por favor, selecione uma opção entre 1 e 4.")
	}
}
