// faça um programa que verifique se o item que a pessoa escolheu comprar na loja está na lista
// lista: laranja, banana, abacaxi, melancia, uva, morango

package main

import "fmt"

func main() {

	var opcao string
	fmt.Println("Entre com o item:")
	fmt.Scanf("%s", &opcao)

	switch opcao {
	case "laranja", "banana", "abacaxi", "melancia", "uva", "morango":
		fmt.Println("Item disponível para compra.")
	default:
		fmt.Println("Item indisponível para compra.")
	}
}
